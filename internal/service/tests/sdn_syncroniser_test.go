package tests

import (
	"context"
	"github.com/stretchr/testify/mock"
	mocks2 "gluck1986/test_work_xml/internal/dataSource/mocks"
	"gluck1986/test_work_xml/internal/model"
	"gluck1986/test_work_xml/internal/service"
	"io"
	"log"
	"testing"
	"time"
)

func TestSdnSyncroniser_SyncroniseIdle(t1 *testing.T) {
	ctx := context.Background()
	parser := mocks2.NewISdnParser(t1)
	stop := make(chan struct{})

	parser.On("Next").Return(func() (model.SdnParseResponse, bool) {
		<-stop
		return model.SdnParseResponse{}, false
	})

	repo := mocks2.NewISdnRepository(t1)

	dep := &service.SdnSyncroniserParams{
		Parser: parser,
		Repo:   repo,
		Log:    log.New(io.Discard, "", 0),
	}
	t := service.NewSdnSyncroniser(dep)
	go func() {
		if err := t.Syncronise(ctx); err != nil {
			t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
		}
	}()
	time.Sleep(time.Millisecond)
	if !t.IsIdle() {
		t1.Errorf("IsIdle = false, want true")
	}
	stop <- struct{}{}
	close(stop)
	time.Sleep(time.Millisecond)
	if t.IsIdle() {
		t1.Errorf("IsIdle = true, want false")
	}
}

func TestSdnSyncroniser_SyncroniseIdleAbort(t1 *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	parser := mocks2.NewISdnParser(t1)

	parser.On("Next").Return(func() (model.SdnParseResponse, bool) {
		time.Sleep(time.Microsecond)
		return model.SdnParseResponse{}, true
	})

	repo := mocks2.NewISdnRepository(t1)
	repo.On("WriteMany", mock.Anything).Return(nil)

	dep := &service.SdnSyncroniserParams{
		Parser: parser,
		Repo:   repo,
		Log:    log.New(io.Discard, "", 0),
	}
	t := service.NewSdnSyncroniser(dep)
	go func() {
		if err := t.Syncronise(ctx); err != nil {
			t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
		}
	}()
	time.Sleep(time.Millisecond)
	if !t.IsIdle() {
		t1.Errorf("IsIdle = false, want true")
	}
	cancel()
	time.Sleep(time.Millisecond)
	if t.IsIdle() {
		t1.Errorf("IsIdle = true, want false")
	}
}

func TestSdnSyncroniser_SyncroniseUseParserOneBatch(t1 *testing.T) {
	ctx := context.Background()
	parser := mocks2.NewISdnParser(t1)
	counter := 0
	maxResp := 500
	parser.On("Next").Return(func() (model.SdnParseResponse, bool) {
		if counter >= maxResp {
			return model.SdnParseResponse{}, false
		}
		counter++
		return model.SdnParseResponse{
			PublishInformation: model.PublishInformation{
				PublishDate: "05/19/2023",
				RecordCount: 0,
			},
		}, true
	}).Times(501)
	batch := make([]model.SdnEntity, 500)
	for i := 0; i < 500; i++ {
		batch[i] = model.SdnEntity{
			UID:         0,
			FirstName:   "",
			LastName:    "",
			PublishDate: time.Date(2023, 5, 19, 0, 0, 0, 0, time.UTC),
		}
	}

	repo := mocks2.NewISdnRepository(t1)
	repo.On("WriteMany", batch).Once().Return(nil)

	dep := &service.SdnSyncroniserParams{
		Parser: parser,
		Repo:   repo,
	}
	t := service.NewSdnSyncroniser(dep)

	if err := t.Syncronise(ctx); err != nil {
		t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
	}
}

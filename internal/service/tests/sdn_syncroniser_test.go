package tests

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gluck1986/test_work_xml/internal/datasource"
	mocks2 "gluck1986/test_work_xml/internal/datasource/mocks"
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

	cache := mocks2.NewIUidCache(t1)

	dep := &service.SdnSyncroniserParams{
		Writer: repo,
		Log:    log.New(io.Discard, "", 0),
		Cache:  cache,
	}
	t := service.NewSdnSyncroniser(dep)
	err := t.Init(ctx, parser)
	if err != nil {
		t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
	}
	go func() {
		if err := t.Syncronise(); err != nil {
			t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
		}
	}()
	time.Sleep(time.Millisecond * 10)
	if !t.IsIdle() {
		t1.Errorf("IsIdle = false, want true")
	}
	stop <- struct{}{}
	close(stop)
	time.Sleep(time.Millisecond * 10)
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

	cache := mocks2.NewIUidCache(t1)
	cache.On("Add", mock.Anything)
	cache.On("Has", mock.Anything).Return(false)

	dep := &service.SdnSyncroniserParams{
		Writer: repo,
		Log:    log.New(io.Discard, "", 0),
		Cache:  cache,
	}
	t := service.NewSdnSyncroniser(dep)
	err := t.Init(ctx, parser)
	if err != nil {
		t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
	}
	go func() {
		if err := t.Syncronise(); err != nil {
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

	cache := mocks2.NewIUidCache(t1)
	cache.On("Add", mock.Anything)
	cache.On("Has", mock.Anything).Return(false)

	dep := &service.SdnSyncroniserParams{
		Writer: repo,
		Cache:  cache,
	}
	t := service.NewSdnSyncroniser(dep)
	err := t.Init(ctx, parser)
	if err != nil {
		t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
	}

	if err := t.Syncronise(); err != nil {
		t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
	}
}

func TestSdnSyncroniser_SyncroniseCache(t1 *testing.T) {
	ctx := context.Background()
	parser := mocks2.NewISdnParser(t1)
	counter := 0
	maxResp := 10
	repeat := true
	parser.On("Next").Return(func() (model.SdnParseResponse, bool) {
		if counter >= maxResp {
			if repeat {
				repeat = false
				counter = 0
			} else {
				return model.SdnParseResponse{}, false
			}
		}
		counter++
		return model.SdnParseResponse{
			Data: model.SdnExternalEntity{
				UID: counter,
			},
			PublishInformation: model.PublishInformation{
				PublishDate: "05/19/2023",
				RecordCount: 0,
			},
		}, true
	}).Times(21)

	repo := mocks2.NewISdnRepository(t1)

	repo.On("WriteMany", mock.MatchedBy(func(input []model.SdnEntity) bool {
		wanted := 10
		given := len(input)
		if given != wanted {
			t1.Errorf("WriteMany receive %v models, wanted count %v", given, wanted)
		}
		return true
	})).Return(nil)

	cache := datasource.NewUidCache()

	dep := &service.SdnSyncroniserParams{
		Writer: repo,
		Cache:  cache,
	}
	t := service.NewSdnSyncroniser(dep)
	err := t.Init(ctx, parser)
	if err != nil {
		t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
	}

	if err := t.Syncronise(); err != nil {
		t1.Errorf("Syncronise() error = %v, wantErr %v", err, false)
	}
}

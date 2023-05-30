package tests

import (
	"github.com/stretchr/testify/mock"
	mocks2 "gluck1986/test_work_xml/internal/datasource/mocks"
	"gluck1986/test_work_xml/internal/service"
	"gluck1986/test_work_xml/internal/service/mocks"
	"log"
	"os"
	"testing"
)

func TestSyncroniseVisor_GetStatusIdle(t1 *testing.T) {

	syncroniserService := mocks.NewISdnSyncroniser(t1)
	repo := mocks2.NewISdnRepository(t1)
	logger := log.New(os.Stdout, "test ", 0)

	syncroniserService.On("IsIdle").Return(true)

	stateService := service.NewSyncroniseVisor(syncroniserService, repo, logger)

	got, err := stateService.GetStatus()
	if err != nil {
		t1.Errorf("GetStatus() error = %v, wantErr %v", err, nil)
		return
	}
	want := service.SyncroniserInProgress
	if got != want {
		t1.Errorf("GetStatus() got = %v, want %v", got, want)
	}
}

func TestSyncroniseVisor_GetStatusEmpty(t1 *testing.T) {

	syncroniserService := mocks.NewISdnSyncroniser(t1)
	repo := mocks2.NewISdnRepository(t1)
	logger := log.New(os.Stdout, "test ", 0)

	syncroniserService.On("IsIdle").Return(false)
	repo.On("Exists", mock.Anything).Return(false, nil)

	stateService := service.NewSyncroniseVisor(syncroniserService, repo, logger)

	got, err := stateService.GetStatus()
	if err != nil {
		t1.Errorf("GetStatus() error = %v, wantErr %v", err, nil)
		return
	}
	want := service.SyncroniserEmpty
	if got != want {
		t1.Errorf("GetStatus() got = %v, want %v", got, want)
	}
}

func TestSyncroniseVisor_GetStatusOk(t1 *testing.T) {

	syncroniserService := mocks.NewISdnSyncroniser(t1)
	repo := mocks2.NewISdnRepository(t1)
	logger := log.New(os.Stdout, "test ", 0)

	syncroniserService.On("IsIdle").Return(false)
	repo.On("Exists", mock.Anything).Return(true, nil)

	stateService := service.NewSyncroniseVisor(syncroniserService, repo, logger)

	got, err := stateService.GetStatus()
	if err != nil {
		t1.Errorf("GetStatus() error = %v, wantErr %v", err, nil)
		return
	}
	want := service.SyncroniserOk
	if got != want {
		t1.Errorf("GetStatus() got = %v, want %v", got, want)
	}
}

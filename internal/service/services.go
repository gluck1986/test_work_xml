package service

import (
	"context"
	"errors"
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/model"
)

// SyncroniserState state of ISdnSyncroniser and its data
type SyncroniserState int

const (
	// SyncroniserEmpty there are no data parsed before and no running parsing now
	SyncroniserEmpty SyncroniserState = iota
	// SyncroniserInProgress parsing in progress
	SyncroniserInProgress
	// SyncroniserOk there are parsed data and no running process in progress
	SyncroniserOk
)

// ErrorSyncroniserAlreadyInProgress will fire if you try run ISdnSyncroniser.Syncronise() or ISdnSyncroniser.Init()
// when  ISdnSyncroniser.Syncronise() is already in progress
var ErrorSyncroniserAlreadyInProgress = errors.New("parsing already in progress")

// ISdnSyncroniser is a syncronise sdn data service from source to storage
//
//go:generate mockery --dir . --name ISdnSyncroniser --output ./mocks
type ISdnSyncroniser interface {
	IsIdle() bool
	Syncronise() error
	Init(ctx context.Context, parser datasource.ISdnParser) error
}

// ISdnWriter write a batch of model.SdnEntity
//
//go:generate mockery --dir . --name ISdnWriter --output ./mocks
type ISdnWriter interface {
	WriteMany([]model.SdnEntity) error
}

// ISyncroniseVisor returns ISdnSyncroniser work status and its data
//
//go:generate mockery --dir . --name ISyncroniseVisor --output ./mocks
type ISyncroniseVisor interface {
	GetStatus() (SyncroniserState, error)
}

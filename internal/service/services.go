package service

import (
	"context"
	"errors"
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/model"
)

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

package service

import (
	"context"
	"gluck1986/test_work_xml/internal/model"
)

// ISdnSyncroniser is a syncronise sdn data service from source to storage
//
//go:generate mockery --dir . --name ISdnSyncroniser --output ./mocks
type ISdnSyncroniser interface {
	IsIdle() bool
	Syncronise(ctx context.Context) error
}

//go:generate mockery --dir . --name ISdnWriter --output ./mocks
type ISdnWriter interface {
	Write(model.SdnEntity) error
}

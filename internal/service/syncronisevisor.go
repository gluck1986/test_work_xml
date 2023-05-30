package service

import (
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/datasource/criteria"
	"log"
)

// SyncroniseVisor checks status ISdnSyncroniser and its data
type SyncroniseVisor struct {
	target ISdnSyncroniser
	store  datasource.ISdnRepository
	log    *log.Logger
}

// NewSyncroniseVisor constructor
func NewSyncroniseVisor(target ISdnSyncroniser, store datasource.ISdnRepository, log *log.Logger) *SyncroniseVisor {
	return &SyncroniseVisor{target: target, store: store, log: log}
}

// GetStatus checks status ISdnSyncroniser and its data
func (t *SyncroniseVisor) GetStatus() (SyncroniserState, error) {
	if t.target.IsIdle() {
		return SyncroniserInProgress, nil
	}
	if exists, err := t.store.Exists(criteria.SdnCriteria{Limit: 1}); err != nil {
		return SyncroniserEmpty, err
	} else if exists {
		return SyncroniserOk, nil
	}
	return SyncroniserEmpty, nil
}

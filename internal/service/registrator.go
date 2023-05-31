package service

import (
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/infrastructure"
)

// Services dependency container
type Services struct {
	SdnSyncroniser  ISdnSyncroniser
	SyncroniseVisor ISyncroniseVisor
}

// NewServices constructor
func NewServices(depSource *datasource.DataSources, depInfr *infrastructure.Infrastructure) *Services {
	syncroniser := NewSdnSyncroniser(&SdnSyncroniserParams{
		Log:    depInfr.Log,
		Writer: depSource.Repositories.SdnRepository,
		Cache:  depSource.UidCache,
	})
	return &Services{
		SdnSyncroniser:  syncroniser,
		SyncroniseVisor: NewSyncroniseVisor(syncroniser, depSource.Repositories.SdnRepository, depInfr.Log),
	}
}

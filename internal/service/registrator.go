package service

import (
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/infrastructure"
)

// Services dependency container
type Services struct {
	SdnSyncroniser ISdnSyncroniser
}

// NewServices constructor
func NewServices(depSource *datasource.DataSources, depInfr *infrastructure.Infrastructure) *Services {
	return &Services{
		SdnSyncroniser: NewSdnSyncroniser(&SdnSyncroniserParams{
			Log:    depInfr.Log,
			Writer: depSource.Repositories.SdnRepository,
		}),
	}
}

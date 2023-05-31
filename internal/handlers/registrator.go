package handlers

import (
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/infrastructure"
	"gluck1986/test_work_xml/internal/service"
)

// Handlers handlers dependency container
type Handlers struct {
	SdnUpdateHandler      *SdnUpdateHandler
	SdnUpdateStateHandler *SdnUpdateStateHandler
	GetNameHandler        *GetNamesHandler
}

type ok struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
	Code   int    `json:"code"`
}

// NewEchoHandlers make handlers
func NewEchoHandlers(srvDep *service.Services, sourceDep *datasource.DataSources, infDep *infrastructure.Infrastructure) *Handlers {
	return &Handlers{
		SdnUpdateHandler:      NewSdnUpdateHandler(srvDep.SdnSyncroniser, sourceDep.ParserFactory, infDep.Log),
		SdnUpdateStateHandler: NewSdnUpdateStateHandler(srvDep.SyncroniseVisor, infDep.Log),
		GetNameHandler:        NewGetNamesHandler(sourceDep.Repositories.SdnRepository, infDep.Log),
	}
}

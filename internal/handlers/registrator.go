package handlers

import (
	"github.com/labstack/echo/v4"
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/infrastructure"
	"gluck1986/test_work_xml/internal/service"
)

// Handlers handlers dependency container
type Handlers struct {
	SdnUpdateHandler echo.HandlerFunc
}

// NewEchoHandlers make handlers
func NewEchoHandlers(srvDep *service.Services, sourceDep *datasource.DataSources, infDep *infrastructure.Infrastructure) *Handlers {
	return &Handlers{
		SdnUpdateHandler: NewExternalUpdateHandler(srvDep.SdnSyncroniser, sourceDep.ParserFactory, infDep.Log),
	}
}

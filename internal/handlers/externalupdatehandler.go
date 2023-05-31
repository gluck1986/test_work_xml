package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/service"
	"log"
	"net/http"
)

// SdnUpdateHandler echo handler
type SdnUpdateHandler struct {
	service       service.ISdnSyncroniser
	parserFactory datasource.ISdnParserFactory
	logger        *log.Logger
}

// NewSdnUpdateHandler constructor
func NewSdnUpdateHandler(
	service service.ISdnSyncroniser,
	factory datasource.ISdnParserFactory,
	logger *log.Logger,
) *SdnUpdateHandler {
	return &SdnUpdateHandler{service: service, parserFactory: factory, logger: logger}
}

// Handle echo handler function
func (t *SdnUpdateHandler) Handle(ctx echo.Context) error {
	backCtx := context.Background()
	parser, err := t.parserFactory.GetParser()
	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, fmt.Errorf("service unavailable"))
	}

	err = t.service.Init(backCtx, parser)
	if err != nil {
		if errors.Is(err, service.ErrorSyncroniserAlreadyInProgress) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
		}
		return echo.NewHTTPError(http.StatusServiceUnavailable, fmt.Errorf("service unavailable"))
	}
	go func() {
		gerr := t.service.Syncronise()
		if gerr != nil {
			t.logger.Println(gerr)
		}
	}()

	okResult := ok{
		Result: true,
		Info:   "",
		Code:   200,
	}

	return ctx.JSON(http.StatusOK, okResult)
}

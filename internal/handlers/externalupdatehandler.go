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

type ok struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
	Code   int    `json:"code"`
}

// NewExternalUpdateHandler constructor
func NewExternalUpdateHandler(
	service service.ISdnSyncroniser,
	factory datasource.ISdnParserFactory,
	logger *log.Logger,
) echo.HandlerFunc {
	handler := &SdnUpdateHandler{service: service, parserFactory: factory, logger: logger}

	return handler.Handle
}

// Handle echo handler function
func (t SdnUpdateHandler) Handle(ctx echo.Context) error {
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

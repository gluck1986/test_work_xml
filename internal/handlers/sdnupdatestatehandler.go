package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gluck1986/test_work_xml/internal/service"
	"log"
	"net/http"
)

// SdnUpdateStateHandler echo handler
type SdnUpdateStateHandler struct {
	service service.ISyncroniseVisor
	logger  *log.Logger
}

// NewSdnUpdateStateHandler constructor
func NewSdnUpdateStateHandler(
	service service.ISyncroniseVisor,
	logger *log.Logger,
) echo.HandlerFunc {
	handler := &SdnUpdateStateHandler{service: service, logger: logger}

	return handler.Handle
}

// Handle echo handler function
func (t *SdnUpdateStateHandler) Handle(ctx echo.Context) error {
	status, err := t.service.GetStatus()

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, fmt.Errorf("service unavailable"))
	}

	result := struct {
		Result bool   `json:"result"`
		Info   string `json:"info"`
	}{}
	switch status {
	case service.SyncroniserInProgress:
		result.Result = false
		result.Info = "updating"
	case service.SyncroniserEmpty:
		result.Result = false
		result.Info = "empty"
	case service.SyncroniserOk:
		result.Result = true
		result.Info = "ok"
	}
	return ctx.JSON(http.StatusOK, result)
}

package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gluck1986/test_work_xml/internal/config"
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/handlers"
	"gluck1986/test_work_xml/internal/infrastructure"
	"gluck1986/test_work_xml/internal/service"
	"gluck1986/test_work_xml/pkg/liberror"
	"net/http"
	"time"
)

// Run api Application
func Run() {
	// config
	cfg := config.Get()

	// logger
	infrastructureDependency := infrastructure.NewInfrastructure(cfg)

	infrastructureDependency.Migrator.Migrate()

	dataSources := datasource.NewSources(infrastructureDependency)

	servicesDependency := service.NewServices(dataSources, infrastructureDependency)

	handlersDependency := handlers.NewEchoHandlers(servicesDependency, dataSources, infrastructureDependency)

	// Initialize Echo instance
	e := echo.New()
	e.HTTPErrorHandler = liberror.Error

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/update", handlersDependency.SdnUpdateHandler)
	e.GET("/state", handlersDependency.SdnUpdateStateHandler)

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal("fatal", "app start server", e.StartServer(s))
}

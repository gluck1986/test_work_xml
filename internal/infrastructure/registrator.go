package infrastructure

import (
	"gluck1986/test_work_xml/internal/config"
	"log"
	"os"
)

// Infrastructure dependency container
type Infrastructure struct {
	Config   *config.Config
	Log      *log.Logger
	Db       *DbContainer
	Migrator *Migrator
}

// NewInfrastructure constructor
func NewInfrastructure(cfg *config.Config) *Infrastructure {
	logger := log.New(os.Stdout, "", 0)
	db := NewDb(cfg, logger)
	return &Infrastructure{
		Config: cfg,
		Log:    logger,
		Db:     db,
		Migrator: NewMigrator(&MigratorParams{
			Db:     db,
			Log:    logger,
			Config: cfg,
		}),
	}
}

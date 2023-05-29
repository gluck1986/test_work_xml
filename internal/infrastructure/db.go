package infrastructure

import (
	"github.com/go-pg/pg/v10"
	"gluck1986/test_work_xml/internal/config"
	"log"
)

type DbContainer struct {
	Db *pg.DB
}

// NewDb db factory
func NewDb(config *config.Config, logger *log.Logger) *DbContainer {
	pgOpts, err := pg.ParseURL(config.PgURL)
	if err != nil {
		logger.Fatal("fatal", "db", err)
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		logger.Fatal("fatal", "db", "first command", err)
	}
	return &DbContainer{Db: pgDB}
}

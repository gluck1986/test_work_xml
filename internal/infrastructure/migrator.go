package infrastructure

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //side effect
	_ "github.com/golang-migrate/migrate/v4/source/file"       //side effect
	"gluck1986/test_work_xml/internal/config"
	"log"
)

// Migrator migration executor
type Migrator struct {
	db     *DbContainer
	log    *log.Logger
	config *config.Config
}

// MigratorParams constructor params
type MigratorParams struct {
	Db     *DbContainer
	Log    *log.Logger
	Config *config.Config
}

// NewMigrator constructor
func NewMigrator(migratorParams *MigratorParams) *Migrator {
	return &Migrator{
		db:     migratorParams.Db,
		log:    migratorParams.Log,
		config: migratorParams.Config,
	}
}

// Migrate migrate database
func (t *Migrator) Migrate() {
	m, err := migrate.New(
		t.config.PgMigrationsPath,
		t.config.PgURL,
	)
	if err != nil {
		t.log.Fatal("fatal ", "migrator init, ", "path:", t.config.PgMigrationsPath, ". ", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		t.log.Fatal("fatal ", "migrator up ", "path:", t.config.PgMigrationsPath, ". ", err)
	}
}

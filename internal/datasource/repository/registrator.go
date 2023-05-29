package repository

import (
	"github.com/go-pg/pg/v10"
	"log"
)

// Repositories dependency container
type Repositories struct {
	SdnRepository *SdnRepository
}

// NewRepositories constructor
func NewRepositories(db *pg.DB, logger *log.Logger) *Repositories {
	return &Repositories{
		SdnRepository: NewSdnRepository(db, logger),
	}
}

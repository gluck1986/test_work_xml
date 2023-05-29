package repository

import (
	"github.com/go-pg/pg/v10"
	"gluck1986/test_work_xml/internal/datasource/criteria"
	"gluck1986/test_work_xml/internal/model"
	"log"
)

// SdnRepository repository to store and read Sdn entities
type SdnRepository struct {
	db     *pg.DB
	logger *log.Logger
}

// NewSdnRepository constructor
func NewSdnRepository(db *pg.DB, logger *log.Logger) *SdnRepository {
	return &SdnRepository{db: db, logger: logger}
}

// Write single entity (create/rewrite)
func (t *SdnRepository) Write(entity model.SdnEntity) error {
	return nil
}

// WriteMany write batch (create/rewrite)
func (t *SdnRepository) WriteMany(models []model.SdnEntity) error {
	t.logger.Println("noop write sdn, count:", len(models))
	return nil
}

// ReadOne read by uid
func (t *SdnRepository) ReadOne(uid int) (model.SdnEntity, error) {
	return model.SdnEntity{}, nil
}

// ReadMany read by criteria many entities
func (t *SdnRepository) ReadMany(criteria criteria.SdnCriteria) ([]model.SdnEntity, error) {
	return nil, nil
}
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
	res, err := t.db.Model(&models).
		OnConflict("(uid) DO UPDATE").
		Set("firstname = EXCLUDED.firstname, lastname = EXCLUDED.lastname, publish = EXCLUDED.publish").
		Insert()
	if err != nil {
		t.logger.Println("error: SdnRepository.WriteMany: ", err)
		return err
	}
	t.logger.Println("debug: SdnRepository.WriteMany; ", "insertions/updates: ", res.RowsAffected())

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

// Exists check are there data by criteria
func (t *SdnRepository) Exists(criteria criteria.SdnCriteria) (bool, error) {
	sdn := new(model.SdnEntity)
	q := t.db.Model(sdn)
	if criteria.Limit > 0 {
		q.Limit(criteria.Limit)
	}
	exists, err := q.Exists()
	if err != nil {
		t.logger.Println("error, SdnRepository, Exists()", err)

		return false, err
	}
	return exists, nil
}

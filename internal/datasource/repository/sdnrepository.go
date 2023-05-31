package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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
func (t *SdnRepository) ReadMany(srcCriteria criteria.SdnCriteria) ([]model.SdnEntity, error) {
	var entities []model.SdnEntity
	q := t.db.Model(&entities)
	switch srcCriteria.Mode {
	case criteria.SdnModeWeak:
		t.applyWeak(q, srcCriteria)
	case criteria.SdnModeStrong:
		t.applyStrong(q, srcCriteria)
	default:
		t.applyWeak(q, srcCriteria)
	}
	if srcCriteria.Limit > 0 {
		q.Limit(srcCriteria.Limit)
	}
	err := q.Select()
	if err != nil {
		t.logger.Println("error, SdnRepository, ReadMany: ", err)
		return nil, err
	}
	return entities, nil
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

func (t *SdnRepository) applyWeak(q *pg.Query, sdnCriteria criteria.SdnCriteria) {
	q.WhereGroup(func(q1 *orm.Query) (*orm.Query, error) {
		{
			fna := sdnCriteria.MaybeFirstName
			if fna != "" {
				q1.WhereOr("firstname ILIKE '%' || ? || '%' OR lastname ILIKE '%' || ? || '%'", fna, fna)
			}
		}
		{
			lna := sdnCriteria.MaybeLastName
			if lna != "" {
				q1.WhereOr("firstname ILIKE '%' || ? || '%' OR lastname ILIKE '%' || ? || '%'", lna, lna)
			}
		}

		return q1, nil
	})
}

func (t *SdnRepository) applyStrong(q *pg.Query, sdnCriteria criteria.SdnCriteria) {
	fna := sdnCriteria.MaybeFirstName
	lna := sdnCriteria.MaybeLastName
	if fna == "" && lna == "" {
		return
	}
	if fna == "" || lna == "" {
		if fna != "" {
			q.Where("firstname ILIKE ? OR lastname ILIKE ?", fna, fna)
		}
		if lna != "" {
			q.Where("firstname ILIKE ? OR lastname ILIKE ?", lna, lna)
		}
		return
	}
	q.Where("firstname ILIKE ? AND lastname ILIKE ?) OR (firstname ILIKE ? AND lastname ILIKE ?", fna, lna, lna, fna)
}

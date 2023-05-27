package dataSource

import (
	"gluck1986/test_work_xml/internal/dataSource/criteria"
	"gluck1986/test_work_xml/internal/model"
)

// ISdnRepository store and read Sdn entities
//
//go:generate mockery --dir . --name ISdnRepository --output ./mocks
type ISdnRepository interface {
	Write(entity model.SdnEntity) error
	WriteMany([]model.SdnEntity) error
	ReadOne(uid int) (model.SdnEntity, error)
	ReadMany(criteria criteria.SdnCriteria)
}

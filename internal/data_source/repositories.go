package data_source

import (
	"gluck1986/test_work_xml/internal/data_source/criteria"
	"gluck1986/test_work_xml/internal/model"
)

//go:generate mockery --dir . --name ISdnRepository --output ./mocks
type ISdnRepository interface {
	Write(entity model.SdnEntity) error
	WriteMany([]model.SdnEntity) error
	ReadOne(uid int) (model.SdnEntity, error)
	ReadMany(criteria criteria.SdnCriteria)
}

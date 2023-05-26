package data_source

import (
	"gluck1986/test_work_xml/internal/model"
	"io"
)

//go:generate mockery --dir . --name ISdnReader --output ./mocks
type ISdnReader interface {
	io.ReadCloser
}

//go:generate mockery --dir . --name ISdnParser --output ./mocks
type ISdnParser interface {
	Next() (model.SdnParseResponse, bool)
}

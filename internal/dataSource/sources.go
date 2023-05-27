package dataSource

import (
	"gluck1986/test_work_xml/internal/model"
	"io"
)

// ISdnReader reader sdn data from anywhere
//
//go:generate mockery --dir . --name ISdnReader --output ./mocks
type ISdnReader interface {
	io.ReadCloser
}

// ISdnParser make temporary entity of Sdn
//
//go:generate mockery --dir . --name ISdnParser --output ./mocks
type ISdnParser interface {
	Next() (model.SdnParseResponse, bool)
}

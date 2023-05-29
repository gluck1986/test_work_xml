package datasource

import (
	"gluck1986/test_work_xml/internal/config"
	"log"
)

// SdnXMLParserFactory sdn parser factory
type SdnXMLParserFactory struct {
	path   string
	logger *log.Logger
}

// NewSdnXMLParserFactory constructor
func NewSdnXMLParserFactory(cfg *config.Config, logger *log.Logger) ISdnParserFactory {
	return &SdnXMLParserFactory{path: cfg.SdnXMLSource, logger: logger}
}

// GetParser open and start read from XML datasource, returns ready to work parser with active data reader
func (t *SdnXMLParserFactory) GetParser() (ISdnParser, error) {
	reader, err := NewSdnHTTPReader(t.path)
	if err != nil {
		return nil, err
	}

	return NewSdnXMLParser(reader), nil
}

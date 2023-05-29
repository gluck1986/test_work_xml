package datasource

import (
	"encoding/xml"
	"gluck1986/test_work_xml/internal/model"
	"strings"
)

const defaultAllowedSdnType = "Individual"

// SdnXMLParser xml parser
type SdnXMLParser struct {
	reader      ISdnReader
	sdnOnlyType string
	decoder     *xml.Decoder
	publishInfo model.PublishInformation
}

// NewSdnXMLParser constructor
func NewSdnXMLParser(reader ISdnReader) ISdnParser {
	return &SdnXMLParser{
		reader:      reader,
		sdnOnlyType: defaultAllowedSdnType,
		decoder:     xml.NewDecoder(reader),
	}
}

// Next returns new model.SdnParseResponse until source has data
func (t *SdnXMLParser) Next() (model.SdnParseResponse, bool) {
	var publishInformation model.PublishInformation
	for {
		token, _ := t.decoder.Token()
		if token == nil {
			err := t.reader.Close()
			if err != nil {
				return model.SdnParseResponse{}, false
			}
			return model.SdnParseResponse{}, false
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "sdnEntry" {
				var sdnExternalEntity model.SdnExternalEntity
				err := t.decoder.DecodeElement(&sdnExternalEntity, &se)
				if err != nil {
					continue
				}
				if strings.TrimSpace(sdnExternalEntity.SdnType) != t.sdnOnlyType {
					continue
				}
				return model.SdnParseResponse{
					Data:               sdnExternalEntity,
					PublishInformation: t.publishInfo,
				}, true
			} else if se.Name.Local == "publshInformation" {
				err := t.decoder.DecodeElement(&publishInformation, &se)
				if err != nil {
					continue
				}
				t.publishInfo = publishInformation
			}
		}
	}
}

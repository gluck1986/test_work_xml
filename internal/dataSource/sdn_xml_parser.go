package dataSource

import (
	"encoding/xml"
	"gluck1986/test_work_xml/internal/model"
	"strings"
)

const DefaultAllowedSdnType = "Individual"

type SdnXmlParser struct {
	reader      ISdnReader
	sdnOnlyType string
	decoder     *xml.Decoder
	publishInfo model.PublishInformation
}

func NewSdnXmlParser(reader ISdnReader) ISdnParser {
	return &SdnXmlParser{
		reader:      reader,
		sdnOnlyType: DefaultAllowedSdnType,
		decoder:     xml.NewDecoder(reader),
	}
}

func (t *SdnXmlParser) Next() (model.SdnParseResponse, bool) {
	var publishInformation model.PublishInformation
	for {
		token, _ := t.decoder.Token()
		if token == nil {
			t.reader.Close()
			return model.SdnParseResponse{}, false
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "sdnEntry" {
				var sdnExternalEntity model.SdnExternalEntity
				t.decoder.DecodeElement(&sdnExternalEntity, &se)
				if strings.TrimSpace(sdnExternalEntity.SdnType) != t.sdnOnlyType {
					continue
				}
				return model.SdnParseResponse{
					Data:               sdnExternalEntity,
					PublishInformation: t.publishInfo,
				}, true
			} else if se.Name.Local == "publshInformation" {
				t.decoder.DecodeElement(&publishInformation, &se)
				t.publishInfo = publishInformation
			}
		}
	}
}

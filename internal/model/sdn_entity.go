package model

import (
	"encoding/xml"
	"time"
)

// SdnParseResponse temporary entity from parser
type SdnParseResponse struct {
	Data               SdnExternalEntity
	PublishInformation PublishInformation
}

// PublishInformation common data for several SdnExternalEntity
type PublishInformation struct {
	XMLName     xml.Name `xml:"publshInformation"`
	PublishDate string   `xml:"Publish_Date"`
	RecordCount int      `xml:"Record_Count"`
}

// SdnExternalEntity Temporary Sdn data
type SdnExternalEntity struct {
	XMLName   xml.Name `xml:"sdnEntry"`
	UID       int      `xml:"uid"`
	FirstName string   `xml:"firstName"`
	LastName  string   `xml:"lastName"`
	SdnType   string   `xml:"sdnType"`
}

// SdnEntity processed sdn data
type SdnEntity struct {
	//lint:ignore U1000 for data mapper
	tableName   struct{}  `pg:"sdn"`
	UID         int       `pg:"uid,pk"`
	FirstName   string    `pg:"firstname"`
	LastName    string    `pg:"lastname"`
	PublishDate time.Time `pg:"publish"`
}

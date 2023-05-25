package model

import (
	"encoding/xml"
	"time"
)

type SdnParseResponse struct {
	Data               SdnExternalEntity
	PublishInformation PublishInformation
}

type PublishInformation struct {
	XMLName     xml.Name `xml:"publshInformation"`
	PublishDate string   `xml:"Publish_Date"`
	RecordCount int      `xml:"Record_Count"`
}

type SdnExternalEntity struct {
	XMLName   xml.Name `xml:"sdnEntry"`
	UID       int      `xml:"uid"`
	FirstName string   `xml:"firstName"`
	LastName  string   `xml:"lastName"`
	SdnType   string   `xml:"sdnType"`
}

type SdnEntity struct {
	UID         int
	FirstName   string
	LastName    string
	PublishDate time.Time
}

package dataSource

import (
	"github.com/stretchr/testify/assert"
	"gluck1986/test_work_xml/internal/model"
	"io"
	"strings"
	"testing"
)

func TestSdnXmlParser_Next(t1 *testing.T) {

	assertResults := func(t *testing.T, expected, actual model.SdnParseResponse) {
		assert.Equal(t, expected.Data.SdnType, actual.Data.SdnType)
		assert.Equal(t, expected.Data.UID, actual.Data.UID)
		assert.Equal(t, expected.Data.LastName, actual.Data.LastName)
		assert.Equal(t, expected.Data.FirstName, actual.Data.FirstName)
		assert.Equal(t, expected.PublishInformation.PublishDate, actual.PublishInformation.PublishDate)
		assert.Equal(t, expected.PublishInformation.RecordCount, actual.PublishInformation.RecordCount)

	}

	reader := io.NopCloser(strings.NewReader(sdnXmlParserTestGetTestXml()))

	parser := NewSdnXmlParser(reader)

	expectedResults := sdnXmlParserTestGetTestResults()

	res, ok := parser.Next()
	assert.True(t1, ok)
	assertResults(t1, expectedResults[0], res)

	res, ok = parser.Next()
	assert.True(t1, ok)
	assertResults(t1, expectedResults[1], res)

	_, ok = parser.Next()
	assert.False(t1, ok)
}

func sdnXmlParserTestGetTestXml() string {
	return `
<?xml version="1.0" standalone="yes"?>
<sdnList xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://tempuri.org/sdnList.xsd">
    <publshInformation>
        <Publish_Date>05/19/2023</Publish_Date>
        <Record_Count>12538</Record_Count>
    </publshInformation>
 <sdnEntry>
        <uid>36</uid>
        <lastName>AEROCARIBBEAN AIRLINES</lastName>
        <sdnType>Entity</sdnType>
        <programList>
            <program>CUBA</program>
        </programList>
        <akaList>
            <aka>
                <uid>12</uid>
                <type>a.k.a.</type>
                <category>strong</category>
                <lastName>AERO-CARIBBEAN</lastName>
            </aka>
        </akaList>
        <addressList>
            <address>
                <uid>25</uid>
                <city>Havana</city>
                <country>Cuba</country>
            </address>
        </addressList>
    </sdnEntry>
	 <sdnEntry>
        <uid>36</uid>
        <lastName>doe</lastName>
        <firstName>john</firstName>
        <sdnType>Individual</sdnType>
    </sdnEntry>
	<sdnEntry>
        <sdnType>Entity</sdnType>
    </sdnEntry>
	 <sdnEntry>
        <uid>37</uid>
        <lastName>doe2</lastName>
        <firstName>john2</firstName>
        <sdnType>Individual</sdnType>
    </sdnEntry>
</sdnList>

`
}

func sdnXmlParserTestGetTestResults() []model.SdnParseResponse {
	return []model.SdnParseResponse{
		{
			Data: model.SdnExternalEntity{
				UID:       36,
				FirstName: "john",
				LastName:  "doe",
				SdnType:   "Individual",
			},
			PublishInformation: model.PublishInformation{
				PublishDate: "05/19/2023",
				RecordCount: 12538,
			},
		},
		{
			Data: model.SdnExternalEntity{
				UID:       37,
				FirstName: "john2",
				LastName:  "doe2",
				SdnType:   "Individual",
			},
			PublishInformation: model.PublishInformation{
				PublishDate: "05/19/2023",
				RecordCount: 12538,
			},
		},
	}
}

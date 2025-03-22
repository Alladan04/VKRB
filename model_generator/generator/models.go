package generator

import "encoding/xml"

type Parameter struct {
	XMLName      xml.Name `xml:"parameter"`
	Type         string   `xml:"type,attr"`
	ID           string   `xml:"id,attr"`
	ShortName    string   `xml:"shortName,attr"`
	DefaultValue string   `xml:"defaultValue,attr,omitempty"`
}

type Rule struct {
	XMLName   xml.Name `xml:"rule"`
	ID        string   `xml:"id,attr"`
	ShortName string   `xml:"shortName,attr"`
	ResultID  string   `xml:"resultId,attr"`
	Relation  string   `xml:"relation,attr"`
	InitID    string   `xml:"initId,attr"`
}

type Parameters struct {
	XMLName    xml.Name    `xml:"parameters"`
	Parameters []Parameter `xml:"parameter"`
}

type Rules struct {
	XMLName xml.Name `xml:"rules"`
	Rules   []Rule   `xml:"rule"`
}

type Class struct {
	XMLName    xml.Name   `xml:"class"`
	ID         string     `xml:"id,attr"`
	ShortName  string     `xml:"shortName,attr"`
	Parameters Parameters `xml:"parameters"`
	Rules      Rules      `xml:"rules"`
}

type Relation struct {
	XMLName      xml.Name `xml:"relation"`
	ID           string   `xml:"id,attr"`
	ShortName    string   `xml:"shortName,attr"`
	InObj        string   `xml:"inObj,attr"`
	OutObj       string   `xml:"outObj,attr"`
	RelationType string   `xml:"relationType,attr"`
	Content      string   `xml:",chardata"`
}

type Relations struct {
	XMLName   xml.Name   `xml:"relations"`
	Relations []Relation `xml:"relation"`
}

type Model struct {
	XMLName   xml.Name  `xml:"model"`
	FormatVer string    `xml:"formatXmlVersion,attr"`
	ID        string    `xml:"id,attr"`
	ShortName string    `xml:"shortName,attr"`
	Desc      string    `xml:"description,attr"`
	Class     Class     `xml:"class"`
	Relations Relations `xml:"relations"`
}

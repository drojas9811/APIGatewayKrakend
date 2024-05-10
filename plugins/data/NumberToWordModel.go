package data

import (
	"encoding/xml"
	"time"
)

const (
	XMLNumber2WordsBase = `<?xml version="1.0" encoding="utf-8"?>
					<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Bod>
					{COMMAND_CONTENT}				
					</soap:Body>
					</soap:Envelope>`
)

type ReqNumber2WordsModelXML struct {
	UbiNum  string   `xml:"ubiNum"`
	XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`
}

func (r *ReqNumber2WordsModelXML) Init(num2word string) {
	r.UbiNum = num2word
}

type RespNumber2WordsModelJson struct {
	Name      string `json:"name"`
	NumString string `json:"num_string"`
	Time      string `json:"time"`
}

func (r *RespNumber2WordsModelJson) Init(NumString string) {
	r.Name = "DRojas9811"
	r.NumString = NumString
	t := time.Now()
	r.Time = t.String()
}

type RespNumber2WordsModelXML struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body   `xml:"Body"`
}
type Body struct{
	NumberToWordsResponse Body2 `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWordsResponse"`
}
type Body2 struct{
	NumberToWordsResult string `xml:"NumberToWordsResult"`
}
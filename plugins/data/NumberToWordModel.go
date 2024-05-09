package data

import (
	"encoding/xml"
	"time"
)

const (
	XMLNumber2WordsBase = `<?xml version="1.0" encoding="utf-8"?>
					<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
					<soap:Body>
					<NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/">
					{COMMAND_CONTENT}
    				</NumberToWords>					
					</soap:Body>
					</soap:Envelope>`
)

type ReqNumber2WordsModelXML struct {
	UbiNum string `xml:"ubiNum"`
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
	XMLName      xml.Name `xml:"Envelope"`
	Number2Words string   `xml:"Body>m:NumberToWordsResponse>m:NumberToWordsResult"`
}

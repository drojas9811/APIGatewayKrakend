package handler

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	log "plugins/config"
	"plugins/data"
	"strconv"
	"strings"
)

func ReqNumberToWords(req *http.Request) (int, string) {
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: REQUEST :: Number2Words :: IN PROCESS")

	num2Word := (*req).URL.Query().Get("number2word")
	aux := (*req).URL.Query()
	aux.Del("number2word")

	var newResponseXml data.ReqNumber2WordsModelXML
	newResponseXml.Init(num2Word)
	xmlCommandContent, err := xml.Marshal(newResponseXml)
	if err != nil {
		log.Logger.Error("GATEWAY :: Plugin Rest2Soap :: HANDLER :: REQUEST :: Number2Words :: XML body cannot be generated.", err)
		return http.StatusInternalServerError, "Internal Server Error"
	}

	soapRequest := strings.ReplaceAll(data.XMLNumber2WordsBase, "{COMMAND_CONTENT}", string(xmlCommandContent))
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: REQUEST :: Number2Words :: BODY TO SEND:", soapRequest)

	(*req).Body = io.NopCloser(bytes.NewReader([]byte(soapRequest)))
	(*req).Header.Add("Content-Type", "text/xml")
	(*req).ContentLength = int64(len(string(soapRequest)))
	(*req).URL.RawQuery = aux.Encode()
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: REQUEST :: Number2Words :: DONE")
	return http.StatusOK, ""
}

func RespNumberToWords(resp *http.Response) (int, string) {
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words :: IN PROCESS")
	finalJsonBodyResponse:=""
	finalStatusCode:= resp.StatusCode
	switch finalStatusCode {
	case http.StatusOK:
		respBodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return http.StatusInternalServerError, "empty body"
		}
		var responseStruct data.RespNumber2WordsModelXML
		err = xml.Unmarshal(respBodyBytes, &responseStruct)
		if err != nil {
			log.Logger.Error("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words :: xml body invalid.", err)
			return http.StatusInternalServerError, "Internal Server Error"
		}

		var outputResponse data.RespNumber2WordsModelJson
		outputResponse.Init(responseStruct.Number2Words)
		jsonBytesResponse, err := json.Marshal(outputResponse)
		if err != nil {
			log.Logger.Error("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words :: json body invalid.", err)
			return http.StatusInternalServerError, "Internal Server Error"
		}
		finalJsonBodyResponse = string(jsonBytesResponse)
		finalStatusCode = http.StatusOK
	default:
		log.Logger.Warning("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words :: An error has come up from the backend.")
		finalJsonBodyResponse = "An error has come up from the backend."
		finalStatusCode =http.StatusInternalServerError
	}

	(*resp).Body = io.NopCloser(bytes.NewReader([]byte(finalJsonBodyResponse)))
	(*resp).Header.Set("Content-Type", "application/json")
	(*resp).Header.Set("Content-Length", strconv.Itoa(len(finalJsonBodyResponse)))
	(*resp).StatusCode = finalStatusCode

	log.Logger.Debug("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words :: BODY ::", finalJsonBodyResponse)
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words :: STATUS CODE ::", finalStatusCode)
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words :: DONE ::")
	return http.StatusOK, ""
}

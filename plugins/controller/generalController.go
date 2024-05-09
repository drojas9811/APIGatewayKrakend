package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	log "plugins/config"
	"plugins/data"
)

func OutputController(w *http.ResponseWriter, resp *http.Response) error {
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: OUTPUT :: IN PROCESS")
	headersOutput := ""
	for k, hs := range (*resp).Header {
		for _, h := range hs {
			(*w).Header().Add(k, h)
			headersOutput += k + ": " + h + " | "
		}
	}
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: OUTPUT :: STATUSCODE:: ", (*resp).StatusCode)
	(*w).WriteHeader((*resp).StatusCode)
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil || len(respBodyBytes) == 0 {
		log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: OUTPUT :: empty body")
		return errors.New("empty body")
	}
	(*resp).Body.Close()

	(*w).Write(respBodyBytes)
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: OUTPUT :: DONE")
	return nil
}

func ValidateErrorMessage(messageResponse string) bool {
	var errorContent data.ErrorJson

	if err := json.Unmarshal([]byte(messageResponse), &errorContent); err != nil {
		log.Logger.Error("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: ERROR :: Output received is NOT in JSON format:: ", messageResponse)
		return false
	} else {
		log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: ERROR MESSAGE VALIDATOR:: Output received is in JSON format:: ", messageResponse)
		return true
	}
}

func ErrorResponseController(w *http.ResponseWriter, statusCode int, messageResponse string) {
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: ERROR RESPONSE::IN PROCESS")

	isJsonResponse := ValidateErrorMessage(messageResponse)

	if isJsonResponse {
		bodyResponse := messageResponse
		(*w).Header().Set("Content-Type", "application/json")
		(*w).WriteHeader(statusCode)
		io.Copy((*w), io.NopCloser(bytes.NewReader([]byte(bodyResponse))))
	} else {
		bodyResponse := `{"message": "` + messageResponse + `"}`
		(*w).Header().Set("Content-Type", "application/json")
		(*w).WriteHeader(statusCode)
		io.Copy((*w), io.NopCloser(bytes.NewReader([]byte(bodyResponse))))
	}
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: ERROR RESPONSE::FINALIZED")

}

package controller

import (
	"net/http"
	"plugins/common"
	log "plugins/config"
	"plugins/handler"
)

func ResponseController(resp *http.Response, command string) (int, string) {
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER:: RESPONSE ::IN PROCESS")

	statusCode, messageResponse := http.StatusOK, ""

	switch command {
	case common.NumberToWords:
		log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: RESPONSE :: Number2Words")
		statusCode, messageResponse = handler.RespNumberToWords(resp)
	default:
		statusCode, messageResponse = http.StatusInternalServerError, "The endpoint doesn't exist."
	}
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: RESPONSE :: DONE")
	return statusCode, messageResponse
}

package controller

import (
	"net/http"
	"plugins/common"
	log "plugins/config"
	"plugins/handler"
)

func RequestController(req *http.Request, command string) (int, string) {
	statusCode, messageResponse := http.StatusOK, ""

	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER:: REQUEST :: IN PROCESS")	
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: REQUEST :: COMMAND ::", command)
	switch command {
	case common.NumberToWords:
		log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: HANDLER :: REQUEST :: Number2Words")
		statusCode, messageResponse = handler.ReqNumberToWords(req)
	default:
		statusCode, messageResponse = http.StatusInternalServerError, "The endpoint '"+req.URL.Path+"' doesn't exist."
	}
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: CONTROLLER :: REQUEST :: DONE")
	return statusCode, messageResponse
}

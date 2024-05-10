package main

import (
	"context"
	log "plugins/config"
	"plugins/controller"
	"net/http"
)



type registerer string

var ClientRegisterer = registerer("GatewayPlugin")
func (r registerer) RegisterClients(f func(name string, handler func(context.Context, map[string]interface{}) (http.Handler, error))) {
	f(string(r)+"-client", r.PluginRest2Soap_Client)
}

func (r registerer) PluginRest2Soap_Client(_ context.Context, extra map[string]interface{}) (http.Handler, error) {
	log.Logger.Info("GATEWAY :: Plugin Rest2Soap :: 'Injected'.")
	command, err := extra["identifier"].(string)
	if !err {
		log.Logger.Warning("GATEWAY:: Plugin Rest2Soap :: An error has occurred while trying to load identifier.")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Logger.Info("GATEWAY:: Plugin Rest2Soap :: 'IN PROCESS'")

		//Creating Request Object
		statusCode, messageResponse := controller.RequestController(req, command)
		if statusCode != http.StatusOK {
			log.Logger.Warning("GATEWAY :: Plugin Rest2Soap :: STATUS CODE:", statusCode, " - MESSAGE RESPONSE:", messageResponse)
			controller.ErrorResponseController(&w, statusCode, messageResponse)
			return
		}

		//Doing HTTP Request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Logger.Warning("GATEWAY :: Plugin Rest2Soap :: STATUS CODE:", http.StatusInternalServerError, " - MESSAGE RESPONSE:", err.Error())
			controller.ErrorResponseController(&w, http.StatusInternalServerError, "Error - http request")
			return
		}

		//Decoding Response
		statusCode, messageResponse = controller.ResponseController(resp,command)
		if statusCode != http.StatusOK {
			log.Logger.Warning("GATEWAY :: Plugin Rest2Soap:: STATUS CODE:", statusCode, " - MESSAGE RESPONSE:", messageResponse)
			controller.ErrorResponseController(&w, http.StatusInternalServerError, messageResponse)
			return
		}

		//update Response Object
		err = controller.OutputController(&w, resp)
		if err != nil {
			log.Logger.Warning("GATEWAY :: Plugin Rest2Soap:: An error has occurred while trying to process the output controller response:", err.Error())
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

	}), nil
}
func main(){}

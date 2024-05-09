package log

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	 "github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("Krakend")

func LoggerInit() {
	var format = logging.MustStringFormatter(
		` %{time:2006/01/02 - 15:04:05.000} %{color}▶ %{level}%{color:reset} %{message}`,
	)

	krakendLogsConfig := getKrakendTelemetryConfig()

	logger := logging.NewLogBackend(os.Stdout, krakendLogsConfig.TelemetryLogs.Prefix, 0)
	loggerFormatter := logging.NewBackendFormatter(logger, format)
	loggerLeveled := logging.AddModuleLevel(loggerFormatter)

	switch strings.ToUpper(krakendLogsConfig.TelemetryLogs.Level) {
	case "DEBUG":
		loggerLeveled.SetLevel(logging.DEBUG, "")
	case "INFO":
		loggerLeveled.SetLevel(logging.INFO, "")
	case "NOTICE":
		loggerLeveled.SetLevel(logging.NOTICE, "")
	case "WARNING":
		loggerLeveled.SetLevel(logging.WARNING, "")
	case "ERROR":
		loggerLeveled.SetLevel(logging.ERROR, "")
	case "CRITICAL":
		loggerLeveled.SetLevel(logging.CRITICAL, "")
	}

	logging.SetBackend(loggerLeveled)
}

type telemetryConfig struct {
	TelemetryLogs struct {
		Level  string `json:"level"`
		Prefix string `json:"prefix"`
	} `json:"telemetry/logging"`
}

func getKrakendTelemetryConfig() *telemetryConfig {
	// Get the content of the krakend.json file
	fileContentBytes := readKrakendConfigFile("krakend.json")

	// Parse the krakend.json contents into a struct
	var serviceConfig struct {
		ExtraConfig telemetryConfig `json:"extra_config"`
	}

	err := json.Unmarshal(fileContentBytes, &serviceConfig)
	if err != nil {
		serviceConfig.ExtraConfig.TelemetryLogs.Level = "INFO"
		serviceConfig.ExtraConfig.TelemetryLogs.Prefix = "[MobiquityBridge]"
	}

	return &serviceConfig.ExtraConfig
}

func readKrakendConfigFile(fileName string) []byte {
	// Load the contents of the krakend.json file
	fileDataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte(`{"extra_config": {"telemetry/logging": {"level": "INFO", "prefix": "[MobiquityBridge]"}}`)
	}
	return fileDataBytes
}

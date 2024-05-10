package data

import (
	"encoding/json"
	"strconv"
	"time"
)

type errorParameter struct {
	ServerError string `json:"serverError"`
}

// ::::::: Model Error ::::::: struct ::::::://
type ErrorJson struct {
	ErrorParameter   errorParameter `json:"errorParameters"`
	ErrorDateTime    string         `json:"errorDateTime"`
	ErrorDescription string         `json:"errorDescription"`
	ErrorCode        string         `json:"errorCode"`
	ErrorCategory    string         `json:"errorCategory"`
}

func (r *ErrorJson) Init(errorDescription string, statusCode int) {
	r.ErrorParameter.ServerError = "ERROR-1"
	r.ErrorDateTime = time.Now().String()
	r.ErrorCode = strconv.Itoa(statusCode)
	r.ErrorDescription = errorDescription
	r.ErrorCategory = ""
}
func (r *ErrorJson) ToString() string {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "Internal Server Error"
	}
	return string(bytes)
}

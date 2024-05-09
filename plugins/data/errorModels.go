package data

import (
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

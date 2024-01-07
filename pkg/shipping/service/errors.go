package service

import "errors"

var (
	ErrCreateRequest     = errors.New("error in create HTTP request")
	ErrPerformRequest    = errors.New("error in perform HTTP request")
	ErrParseResponseBody = errors.New("error in parse body data")
	ErrInvalidRequest    = errors.New("invalid request")
	ErrMarshal           = errors.New("error in serialize to JSON")
	ErrUnmarshal         = errors.New("error in unmarshal JSON")
)

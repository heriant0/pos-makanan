package infraerror

import "errors"

var (
	ErrorRepository = errors.New("error repository")
	UnknownError    = errors.New("unknown error")
	InvalidRole     = errors.New("invalid role")
)

package handlers

import "errors"

var (
	errRequestBodyInvalid = errors.New("invalid request body")
	errRequestHeaderInvalid = errors.New("invalid request header")
	errRequestSchema = errors.New("invalid request schema input")
)

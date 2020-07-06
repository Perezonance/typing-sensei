package handlers

import (
	log "github.com/perezonance/typing-sensei/sensei-light-logger/pkg/logger"
	"net/http"
)

func writeRes(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	res := []byte(message)
	_, err := w.Write(res)
	if err != nil {
		log.ErrorLog("Failed to write response", err)
	}
	return
}
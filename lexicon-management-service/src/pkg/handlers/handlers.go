package handlers

import (
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/server"
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/storage"

	"net/http"
)

var (
	conf HandlerConfig
	corrID string
)

const (
	headerCorrId = "X-Correlation-ID"
	headerReqId = "X-Request-ID"
)

func init () {
	//Load the handler config...
	c, err := NewConfig()
	if err != nil {
		//TODO: Proper Error Handling
	}
	conf = c
}

func PublicLexHandler(w http.ResponseWriter, r *http.Request) {
	dynamoConf, err := storage.NewDynamoConfig(conf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	serverConf, err := server.NewServerConfig(conf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	db, err := storage.NewDynamo(dynamoConf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	s, err := server.NewServer(db, serverConf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	switch r.Method {
	case http.MethodGet:
		s.GetPubLexicons(w, r, corrID)
	case http.MethodDelete:

	case http.MethodPost:

	case http.MethodPatch:

	}
}

func UserLexHandler(w http.ResponseWriter, r *http.Request) {

}

//Determines if the request header contains a correlation id if empty then generate a new one
func getCorrelationID(r *http.Request) (string, error) {
	headVal := r.Header.Get(http.CanonicalHeaderKey(headerCorrId))

	if headVal == "" {

	}

	return "", nil
}

func getRequestID(r *http.Request) (string, error) {
	r.Header.Get(http.CanonicalHeaderKey(headerReqId))
	return "", nil
}

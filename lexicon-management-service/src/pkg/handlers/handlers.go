package handlers

import (
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/server"
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/storage"

	guuid "github.com/google/uuid"

	"net/http"
)

var (
	conf 	Config
	corrID 	string
	reqID 	string
)

const (
	headerCorrId 	= "X-Correlation-ID"
	headerReqId 	= "X-Request-ID"
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
	corrID = getCorrelationID(r)
	reqID = getRequestID(r)

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

	serverReq, err := parseRequest(r)
	if err != nil {
		//TODO: Proper Error Handling
	}

	switch r.Method {
	case http.MethodGet:
		s.GetPubLexicons(serverReq)
	case http.MethodDelete:
		s.AddPubLexicons(serverReq)
	case http.MethodPost:
		s.DeletePubLexicons(serverReq)
	case http.MethodPatch:
		s.UpdatePubLexicons(serverReq)
	default:
		err := writeRes(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		if err != nil {
			//TODO: Proper Error Handling
		}
	}
	return
}

func UserLexHandler(w http.ResponseWriter, r *http.Request) {

}

//Determines if the request header contains a correlation id if empty then generate a new one
func getCorrelationID(r *http.Request) string {
	headVal := r.Header.Get(http.CanonicalHeaderKey(headerCorrId))

	if headVal == "" {
		return guuid.New().String()
	}
	return headVal
}

func getRequestID(r *http.Request) string {
	headVal := r.Header.Get(http.CanonicalHeaderKey(headerReqId))

	if headVal == "" {
		return guuid.New().String()
	}
	return headVal
}

func parseRequest(r *http.Request) (ServerRequest, error) {
	//TODO: UNIMPLEMENTED
	return ServerRequest{}, nil
}

func writeRes(w http.ResponseWriter, code int, message string) error {
	//TODO: UNIMPLEMENTED
	return nil
}
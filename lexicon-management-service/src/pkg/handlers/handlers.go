package handlers

import (
	"encoding/json"
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/server"
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/storage"
	log "github.com/perezonance/typing-sensei/sensei-light-logger/pkg/logger"

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
		log.ErrorLog("failed to load general configuration", err)
	}
	conf = c
}

func PublicLexHandler(w http.ResponseWriter, r *http.Request) {
	log.InfoLog("Entered")
	corrID = getCorrelationID(r)
	reqID = getRequestID(r)

	dynamoConf, err := storage.NewDynamoConfig(conf)
	if err != nil {
		log.ErrorLog("failed to load dynamo configuration", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	serverConf, err := server.NewServerConfig(conf)
	if err != nil {
		log.ErrorLog("failed to load server configuration", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	db, err := storage.NewDynamo(dynamoConf)
	if err != nil {
		log.ErrorLog("failed to establish a new connection with dynamo", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	s, err := server.NewServer(db, serverConf)
	if err != nil {
		log.ErrorLog("failed to create a new server", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	serverReq, err := parseRequest(r)
	if err != nil {
		log.ErrorLog("failed to parse request", err)
		writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		return
	}

	var res ServerResponse

	switch r.Method {
	case http.MethodGet:
		res, err = s.GetPubLexicons(serverReq)
		if err != nil {
			log.ErrorLog("server failed to getPubLexicons", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	case http.MethodDelete:
		res, err = s.AddPubLexicons(serverReq)
		if err != nil {
			log.ErrorLog("server failed to AddPublicLexicons", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	case http.MethodPost:
		res, err = s.DeletePubLexicons(serverReq)
		if err != nil {
			log.ErrorLog("server failed to DeletePublicLexicons", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	case http.MethodPatch:
		res, err = s.UpdatePubLexicons(serverReq)
		if err != nil {
			log.ErrorLog("server failed to UpdatePublicLexicons", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	default:
		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}

	bod, err := parseServerResponse(res)
	if err != nil {
		log.ErrorLog("failed to parse server response body", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	writeRes(http.StatusOK, bod, w)
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

func parseServerResponse(res ServerResponse) (string, error) {
	bod, err := json.Marshal(res)
	if err != nil {
		log.ErrorLog("failed to marshal server response to []byte", err)
	}

	return string(bod), err
}

func authRequest() {

}
package handlers

import (


	"net/http"
)

func PublicLexHandler(w http.ResponseWriter, r *http.Request) {
	dynamoConf, err := NewDynamoConf(conf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	serverConf, err := NewServerConf(conf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	db, err := NewDynamo(dynamoConf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	s, err := NewServer(db, serverConf)
	if err != nil {
		//TODO: Proper Error Handling
	}

	switch r.Method {
	case http.MethodGet:
		s.GetPubLexicons(w, r, corrId)
	case http.MethodDelete:

	case http.MethodPost:

	case http.MethodPatch:

	}
}

func UserLexHandler(w http.ResponseWriter, r *http.Request) {

}

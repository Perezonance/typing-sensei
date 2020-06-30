package handlers

import (
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/entities"
)

type (

	ServerRequest struct {
		CorrelationID 		string
		RequestID 			string
		Lexicons 			[]entities.Lexicon
	}

	ServerResponse struct {
		CorrelationID 		string
		RequestID 			string
		Lexicons 			[]entities.Lexicon
		HttpStatusCode 		int
		HttpMessage 		string
	}
)

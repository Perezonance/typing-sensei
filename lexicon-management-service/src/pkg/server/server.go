package server

import (
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/handlers"
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/storage"

	"net/http"
)

type (
	Server struct {

	}
)

func NewServer(db *storage.Dynamo, c ServerConfig) (*Server, error) {
	//TODO: UNIMPLEMENTED
	return &Server{}, nil
}

func (s *Server) GetPubLexicons(request handlers.ServerRequest) (handlers.ServerResponse, error){

}

func (s *Server) AddPubLexicons(in handlers.ServerRequest) (handlers.ServerResponse, error){

}

package server

import (
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/handlers"
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/storage"
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
	//TODO: UNIMPLEMENTED
	return handlers.ServerResponse{}, nil
}

func (s *Server) AddPubLexicons(in handlers.ServerRequest) (handlers.ServerResponse, error){
	//TODO: UNIMPLEMENTED
	return handlers.ServerResponse{}, nil
}

func (s *Server) DeletePubLexicons(in handlers.ServerRequest) (handlers.ServerResponse, error) {
	//TODO: UNIMPLEMENTED
	return handlers.ServerResponse{}, nil
}

func (s *Server) UpdatePubLexicons(in handlers.ServerRequest) (handlers.ServerResponse, error) {
	//TODO: UNIMPLEMENTED
	return handlers.ServerResponse{}, nil
}

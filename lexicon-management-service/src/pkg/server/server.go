package server

import (
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

func (s *Server) GetPubLexicons(w http.ResponseWriter, r *http.Request, corrID string) {

}

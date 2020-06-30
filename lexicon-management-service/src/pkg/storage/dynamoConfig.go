package storage

import (
	"github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/handlers"
)

type (
	DynamoConfig struct {

	}
)

func NewDynamoConfig(c handlers.Config) (DynamoConfig, error) {
	//TODO: UNIMPLEMENTED
	return DynamoConfig{}, nil
}

package entities

import "github.com/perezonance/typing-sensei/lexicon-management-service/src/pkg/entities"

type TestTemplate struct {
	TempID 			string 				`json:"temp_id"`
	TempName 		string				`json:"temp_name"`
	Lexicon 		entities.Lexicon	`json:"lexicon"`
	TempTimeLen 	int64				`json:"temp_time_len"`
	TempWordLen 	int64				`json:"temp_word_len"`
	PersistOrder 	bool 				`json:"persist_order"`
	WordList		[]string 			`json:"word_list"`
}

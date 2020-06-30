package entities

type Lexicon struct {
	LexID 			string				`json:"lex_id"`
	UserID 			string 				`json:"user_id"`
	IsPublic 		bool 				`json:"is_public"`
	LexRating 		float32 			`json:"lex_rating"`
	CreationDate 	string 				`json:"creation_date"`
	WordToRatioMap 	map[string]float32	`json:"word_to_ratio_map"`
}


package sentence

import "wumpus-birthday/pkg/globals"

const insertSentence = `
INSERT INTO birthday_sentences(server_id, raw_sentence, sentence_lower) 
VALUES ($1, $2, LOWER($2));`

const removeSentence = `
DELETE FROM birthday_sentences
WHERE server_id = $1 AND sentence_lower = lower($2)
`

const randomGet = `
SELECT raw_sentence FROM birthday_sentences 
WHERE server_id = $1 
ORDER BY RANDOM() LIMIT 1
`

type Sentence struct {
	RawSentence string `db:"raw_sentence,required"`
}

func Add(serverID string, sentence string) error {
	_, err := globals.DB.Exec(insertSentence, serverID, sentence)
	return err
}

func Remove(serverID string, sentence string) error {
	_, err := globals.DB.Exec(removeSentence, serverID, sentence)
	return err
}

func Random(serverID string) (Sentence, error) {
	var sentence = Sentence{}
	err := globals.DB.Select(&sentence, randomGet, serverID)
	return sentence, err
}

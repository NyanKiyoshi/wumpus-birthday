package sentence

import "wumpus-birthday/pkg/globals"

const insertSentence = `
INSERT INTO birthday_sentences(server_id, raw_sentence, sentence_lower) 
VALUES ($1, $2, LOWER($2));`

const removeSentence = `
DELETE FROM birthday_sentences
WHERE server_id = $1 AND sentence_lower = lower($2)
`

func Add(serverID string, sentence string) error {
	_, err := globals.DB.Exec(insertSentence, serverID, sentence)
	return err
}

func Remove(serverID string, sentence string) error {
	_, err := globals.DB.Exec(removeSentence, serverID, sentence)
	return err
}

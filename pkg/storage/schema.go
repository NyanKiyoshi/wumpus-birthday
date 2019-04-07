package storage

// FIXME: add check of uniqueness lower(text)
// Schema defines the schema of the models
var Schema = `
CREATE TABLE IF NOT EXISTS birthdays (
    user_id TEXT NOT NULL,
	server_id TEXT NOT NULL,
    date DATETIME NOT NULL,

	PRIMARY KEY (user_id, server_id)
);

CREATE TABLE IF NOT EXISTS birthday_sentences (
    server_id TEXT NOT NULL,
    raw_sentence TEXT NOT NULL,
    sentence_lower TEXT NOT NULL CHECK (
        sentence_lower = lower(raw_sentence)),
	
    PRIMARY KEY (server_id, sentence_lower)
);
`

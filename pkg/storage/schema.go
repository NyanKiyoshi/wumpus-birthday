package storage

// Schema defines the schema of the models
var Schema = `
CREATE TABLE IF NOT EXISTS birthdays (
    user_id text,
	server_id text,
    date DATETIME,

	PRIMARY KEY (user_id, server_id)
);
`

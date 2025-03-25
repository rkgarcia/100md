package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS questions (
	id    INTEGER PRIMARY KEY,
    question VARCHAR(80)  DEFAULT ''
);

CREATE TABLE IF NOT EXISTS answers (
	id    INTEGER PRIMARY KEY,
    description VARCHAR(80)  DEFAULT '',
    votes INTEGER  DEFAULT 0,
	question_id INTEGER REFERENCES questions(id)
);
`

func getDB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", "data.sqlite3")
	db.MustExec(schema)
	return db
}

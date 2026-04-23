package database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Configure(db *sql.DB) error {
	pragmas := []string{
		"PRAGMA journal_mode = WAL;",
		"PRAGMA foreign_keys = ON;",
	}

	for _, p := range pragmas {
		if _, err := db.Exec(p); err != nil {
			return err
		}
	}

	return nil
}

func InitSchema(db *sql.DB) error {
	schemaQueries := []string{
		`
			CREATE TABLE IF NOT EXISTS USER (
				id TEXT NOT NULL CHECK(length(id) == 36),
				email TEXT NOT NULL UNIQUE CHECK(length(email) <= 255),
				name TEXT NOT NULL CHECK(length(id) == 36),
				password_hash TEXT NOT NULL CHECK(length(password_hash) <= 60),
				ctime INTEGER NOT NULL,
				mtime INTEGER,
				PRIMARY KEY (id)
			) STRICT;
		 `,
		`
		 	REPLACE INTO USER (id, email, name, password_hash, ctime, mtime) VALUES("019dba5d-83a4-7f97-bdf1-97a5fb3d5869", "foo@ba.r", "John Doe", "$2y$10$vskPu8bji.zmwDI9Jvi3C.uXG8IDytaeHLakb4nOzlFyoPTlwBklW", 1776948241, NULL);
		 `,
	}

	for _, p := range schemaQueries {
		if _, err := db.Exec(p); err != nil {
			return err
		}
	}

	return nil
}

func Open(createSchema bool) (*sql.DB, error) {
	databaseDir := filepath.Join(".", "data")
	if err := os.MkdirAll(databaseDir, os.ModePerm); err != nil {
		return nil, err
	}

	databasePath := filepath.Join(databaseDir, "gotask.sqlite3")

	db, err := sql.Open("sqlite", "file:"+databasePath)
	if err != nil {
		return nil, err
	}

	err = Configure(db)

	if err != nil {
		return nil, err
	}

	if createSchema == true {
		err = InitSchema(db)

		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func Close(db *sql.DB) error {
	return db.Close()
}

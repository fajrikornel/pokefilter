package model

import "database/sql"
import _ "github.com/mattn/go-sqlite3"

type InMemorySqlStore struct {
	db *sql.DB
}

func NewInMemorySqlStore() (*InMemorySqlStore, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	sqlStore := &InMemorySqlStore{db: db}

	return sqlStore, nil
}

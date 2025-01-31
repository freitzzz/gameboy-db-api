package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type database struct {
	*sql.DB
}

func Open(filepath string) (*database, error) {
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", filepath))
	if err != nil {
		return nil, err
	}

	return &database{DB: db}, err
}

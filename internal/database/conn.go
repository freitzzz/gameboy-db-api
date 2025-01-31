package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Open(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", filepath))
	return db, err
}

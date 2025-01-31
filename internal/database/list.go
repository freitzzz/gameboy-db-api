package database

import "database/sql"

func Tables(db *sql.DB) ([]string, error) {
	cursor, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return nil, err
	}

	tables := []string{}
	for cursor.Next() {
		tables = append(tables, "")
		if err = cursor.Scan(&tables[len(tables)-1]); err != nil {
			return nil, err
		}
	}

	return tables, nil
}

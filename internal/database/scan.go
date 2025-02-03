package database

import (
	"database/sql"
	"fmt"
	"reflect"
)

// Scan a database row to a struct using reflection
// Input struct must be a pointer.
func ScanTo(row *sql.Row, to any) error {
	ptrValue := reflect.ValueOf(to)
	if ptrValue.Kind() != reflect.Ptr || ptrValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("to must be a pointer to a struct")
	}

	rValue := ptrValue.Elem()
	args := []any{}
	for i := 0; i < rValue.NumField(); i++ {
		fieldPtr := rValue.Field(i).Addr().Interface()
		args = append(args, fieldPtr)
	}

	err := row.Scan(args...)
	return err
}

// Same as [ScanTo] but returns a slice of the record struct.
func ScanAllTo[T any](rows *sql.Rows) ([]T, error) {
	var t T
	to := []T{}

	ptrValue := reflect.ValueOf(&t)
	rValue := ptrValue.Elem()

	args := []any{}
	for i := 0; i < rValue.NumField(); i++ {
		fieldPtr := rValue.Field(i).Addr().Interface()
		args = append(args, fieldPtr)
	}

	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}

		to = append(to, t)
	}

	return to, nil
}

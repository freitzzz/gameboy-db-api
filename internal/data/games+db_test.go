//go:build integration
// +build integration

package data_test

import (
	"database/sql"
	"testing"
	"testing/quick"

	"github.com/freitzzz/gameboy-db-api/internal/data"
	"github.com/freitzzz/gameboy-db-api/internal/errors"
	"github.com/freitzzz/gameboy-db-api/internal/model"
)

func TestFindReturnsErrRecordNotFoundIfNoRowsAreReturnedFromDatabase(t *testing.T) {
	db, err := sql.Open("sqlite", "file:../../database/db.sqlite")
	if err != nil {
		t.Errorf("did not expect open call to fail, %v", err)
	}
	t.Cleanup(func() { db.Close() })

	r := data.NewDbGamesRepository(db)
	_, err = r.Find(-1)
	if err != errors.ErrRecordNotFound {
		t.Errorf("expected %v but got %v", errors.ErrRecordNotFound, err)
	}
}

func TestPreviewsSearchQueryHandlesOddCharacters(t *testing.T) {
	db, err := sql.Open("sqlite", "file:../../database/db.sqlite")
	if err != nil {
		t.Errorf("did not expect open call to fail, %v", err)
	}
	t.Cleanup(func() { db.Close() })

	r := data.NewDbGamesRepository(db)

	if err := quick.Check(func(x string) bool {
		_, err := r.Previews(model.QueryOptions{Count: 20, Page: 1, Name: x})
		return err == nil
	}, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error(err)
	}
}

func BenchmarkPreviews(b *testing.B) {
	db, err := sql.Open("sqlite", "file:../../database/db.sqlite")
	if err != nil {
		b.Errorf("did not expect open call to fail, %v", err)
	}
	b.Cleanup(func() { db.Close() })

	r := data.NewDbGamesRepository(db)
	for i := 0; i < b.N; i++ {
		r.Previews(model.QueryOptions{Count: 100})
	}
}

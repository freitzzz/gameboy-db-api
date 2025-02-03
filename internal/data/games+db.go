package data

import (
	"database/sql"

	"github.com/freitzzz/gameboy-db-api/internal/database"
	"github.com/freitzzz/gameboy-db-api/internal/errors"
	"github.com/freitzzz/gameboy-db-api/internal/model"
)

type dbGamesRepository struct {
	*sql.DB
}

func NewDbGamesRepository(db *sql.DB) GamesRepository {
	return dbGamesRepository{db}
}

func (r dbGamesRepository) Find(id int) (model.Game, error) {
	var result gameRecord

	cursor := r.QueryRow("SELECT * FROM GameDetails WHERE gameid = ?", id)
	err := database.ScanTo(cursor, &result)
	if err == nil {
		return result.Model(), nil
	}

	if err == sql.ErrNoRows {
		err = errors.ErrRecordNotFound
	}

	return model.Game{}, err
}

func (r dbGamesRepository) Previews(opt model.QueryOptions) ([]model.GamePreview, error) {
	cursor, err := r.Query(gamePreviewQuery(opt), opt.Count)
	if err != nil {
		return nil, err
	}

	result, err := database.ScanAllTo[gamePreviewRecord](cursor)
	if err != nil {
		return nil, err
	}

	gamePreviews := make([]model.GamePreview, len(result))
	for i := range gamePreviews {
		gamePreviews[i] = result[i].Model()
	}

	return gamePreviews, nil
}

func gamePreviewQuery(opt model.QueryOptions) string {
	if opt.Order.HighestRating() {
		return "SELECT * FROM HighestRatedGamePreview LIMIT ?"
	}

	if opt.Order.LowestRating() {
		return "SELECT * FROM LowestRatedGamePreview LIMIT ?"
	}

	return "SELECT * FROM GamePreview LIMIT ?"
}

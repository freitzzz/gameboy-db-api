package data

import "github.com/freitzzz/gameboy-db-api/internal/model"

type GamesRepository interface {
	Previews(opt model.QueryOptions) ([]model.GamePreview, error)
	Find(id int) (model.Game, error)
}

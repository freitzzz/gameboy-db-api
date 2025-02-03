package service

import (
	"github.com/freitzzz/gameboy-db-api/internal/data"
	"github.com/freitzzz/gameboy-db-api/internal/model"
)

const (
	previewListingCount = 20
)

type GamesService struct {
	repository        data.GamesRepository
	cacheHighestRated []model.GamePreview
	cacheLowestRated  []model.GamePreview
	cacheGames        map[int]model.Game
}

func NewGamesService(repository data.GamesRepository) *GamesService {
	return &GamesService{repository: repository, cacheGames: map[int]model.Game{}}
}

func (s *GamesService) Find(id int) (model.Game, error) {
	if game, ok := s.cacheGames[id]; ok {
		return game, nil
	}

	game, err := s.repository.Find(id)
	if err != nil {
		return game, err
	}

	s.cacheGames[id] = game
	return game, nil
}

func (s *GamesService) HighestRated() ([]model.GamePreview, error) {
	if s.cacheHighestRated != nil {
		return s.cacheHighestRated, nil
	}

	previews, err := s.repository.Previews(model.QueryOptions{Count: previewListingCount, Order: model.QueryOrderRatingDesc})
	if err != nil {
		return nil, err
	}

	s.cacheHighestRated = previews
	return previews, nil
}

func (s *GamesService) LowestRated() ([]model.GamePreview, error) {
	if s.cacheLowestRated != nil {
		return s.cacheLowestRated, nil
	}

	previews, err := s.repository.Previews(model.QueryOptions{Count: previewListingCount, Order: model.QueryOrderRatingAsc})
	if err != nil {
		return nil, err
	}

	s.cacheLowestRated = previews
	return previews, nil
}

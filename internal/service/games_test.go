package service_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/freitzzz/gameboy-db-api/internal/model"
	"github.com/freitzzz/gameboy-db-api/internal/service"
)

type randomResultsGamesRepository struct{}

func (r randomResultsGamesRepository) Previews(model.QueryOptions) ([]model.GamePreview, error) {
	return make([]model.GamePreview, rand.Intn(5)), nil
}

func (r randomResultsGamesRepository) Find(id int) (model.Game, error) {
	return model.Game{ID: rand.Intn(2000)}, nil
}

func TestListingResultsAreCached(t *testing.T) {
	s := service.NewGamesService(randomResultsGamesRepository{})

	testCases := []struct {
		desc string
		call func() ([]model.GamePreview, error)
	}{
		{
			desc: "calling HighestRated caches results on service",
			call: s.HighestRated,
		},
		{
			desc: "calling LowestedRated caches results on service",
			call: s.LowestRated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r1, err := tC.call()
			if err != nil {
				t.Errorf("did not expect call to fail, %v", err)
			}

			r2, err := tC.call()
			if err != nil {
				t.Errorf("did not expect call to fail, %v", err)
			}

			eq := slices.EqualFunc(r1, r2, func(e1, e2 model.GamePreview) bool {
				return e1.ID == e2.ID
			})

			if !eq {
				t.Errorf("expected call results to be cached, got %v and %v", r1, r2)
			}
		})
	}
}

func TestFindResultIsCached(t *testing.T) {
	s := service.NewGamesService(randomResultsGamesRepository{})
	id := 1337

	g1, err := s.Find(id)
	if err != nil {
		t.Errorf("did not expect call to fail, %v", err)
	}

	g2, err := s.Find(id)
	if err != nil {
		t.Errorf("did not expect call to fail, %v", err)
	}

	if g1.ID != g2.ID {
		t.Errorf("expected call results to be cached, got %v, %v", g1, g2)
	}
}

func TestSearchResultsAreCached(t *testing.T) {
	s := service.NewGamesService(randomResultsGamesRepository{})

	testCases := []struct {
		desc string
		call func() ([]model.GamePreview, error)
	}{
		{
			desc: "calling Search caches results on service",
			call: func() ([]model.GamePreview, error) { return s.Search(1, "pokemon") },
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r1, err := tC.call()
			if err != nil {
				t.Errorf("did not expect call to fail, %v", err)
			}

			r2, err := tC.call()
			if err != nil {
				t.Errorf("did not expect call to fail, %v", err)
			}

			eq := slices.EqualFunc(r1, r2, func(e1, e2 model.GamePreview) bool {
				return e1.ID == e2.ID
			})

			if !eq {
				t.Errorf("expected call results to be cached, got %v and %v", r1, r2)
			}
		})
	}
}

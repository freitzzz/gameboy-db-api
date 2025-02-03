package data

import (
	"strings"

	"github.com/freitzzz/gameboy-db-api/internal/model"
)

type gameRecord struct {
	ID              *int
	Name            *string
	Description     *string
	ReleaseYear     *int
	ESRB            *int
	Trivia          *string
	Promo           *string
	Adult           *bool
	Rating          *int
	Critics         *int
	Genres          *string
	Platforms       *string
	Developers      *string
	Publishers      *string
	Screenshots     *string
	ScreenshotsHash *string
	Thumbnail       *string
	ThumbnailHash   *string
	Cover           *string
	CoverHash       *string
	Gameplay        *string
}

type gamePreviewRecord struct {
	ID            *int
	Name          *string
	Genres        *string
	Platforms     *string
	ThumbnailUrl  *string
	ThumbnailHash *string
}

func (g gameRecord) Model() model.Game {
	return model.Game{
		ID:          *g.ID,
		Name:        *g.Name,
		Description: g.Description,
		ReleaseYear: *g.ReleaseYear,
		ESRB:        *g.ESRB,
		Adult:       *g.Adult,
		Promo:       g.Promo,
		Trivia:      g.Trivia,
		Rating:      g.Rating,
		Critics:     g.Critics,
		Thumbnail:   assetColumns(g.Thumbnail, g.ThumbnailHash),
		Cover:       assetColumns(g.Cover, g.CoverHash),
		Screenshots: sliceAssetColumns(g.Screenshots, g.ScreenshotsHash),
		Genres:      splitColumnValues(g.Genres),
		Platforms:   splitColumnValues(g.Genres),
		Developers:  splitColumnValues(g.Genres),
		Publishers:  splitColumnValues(g.Genres),
	}
}

func (g gamePreviewRecord) Model() model.GamePreview {
	return model.GamePreview{
		ID:        *g.ID,
		Name:      *g.Name,
		Genres:    splitColumnValues(g.Genres),
		Platforms: splitColumnValues(g.Platforms),
		Thumbnail: assetColumns(g.ThumbnailUrl, g.ThumbnailHash),
	}
}

func assetColumns(url *string, previewHash *string) *model.GameAsset {
	if url == nil {
		return nil
	}

	return &model.GameAsset{URL: *url, PreviewHash: previewHash}
}

func sliceAssetColumns(urls *string, previewHashes *string) []model.GameAsset {
	if urls == nil {
		return nil
	}

	urlsSplit := splitColumnValues(urls)
	previewHashesSplit := splitColumnValues(previewHashes)
	assets := make([]model.GameAsset, len(urlsSplit))
	for i := range urlsSplit {
		var hash *string
		if previewHashesSplit != nil && len(previewHashesSplit) < i {
			hash = &previewHashesSplit[i]
		}

		assets[i] = model.GameAsset{URL: urlsSplit[i], PreviewHash: hash}
	}

	return assets
}

func splitColumnValues(column *string) []string {
	if column == nil {
		return nil
	}

	return strings.Split(*column, ", ")
}

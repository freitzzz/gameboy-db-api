package model

type Game struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	ReleaseYear int         `json:"releaseYear"`
	ESRB        int         `json:"esrb"`
	Genres      []string    `json:"genres"`
	Developers  []string    `json:"developers"`
	Publishers  []string    `json:"publishers"`
	Platforms   []string    `json:"platforms"`
	Screenshots []GameAsset `json:"screenshots"`
	Adult       bool        `json:"adult"`
	Description *string     `json:"description,omitempty"`
	Promo       *string     `json:"promo,omitempty"`
	Trivia      *string     `json:"trivia,omitempty"`
	Rating      *int        `json:"rating,omitempty"`
	Critics     *int        `json:"critics,omitempty"`
	Thumbnail   *GameAsset  `json:"thumbnail,omitempty"`
	Cover       *GameAsset  `json:"cover,omitempty"`
	Gameplay    *GameAsset  `json:"gameplay,omitempty"`
}

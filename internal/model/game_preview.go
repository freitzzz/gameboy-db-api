package model

type GamePreview struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Genres    []string   `json:"genres"`
	Platforms []string   `json:"platforms"`
	Thumbnail *GameAsset `json:"thumbnail"`
}

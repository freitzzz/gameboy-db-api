package model

type GameAsset struct {
	URL         string  `json:"url"`
	PreviewHash *string `json:"previewHash,omitempty"`
}

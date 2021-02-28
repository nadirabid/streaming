package models

type Video struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	AssetPath   string `json:"-"`

	// Actors??
}

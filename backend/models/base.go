package models

import "time"

// Model contains the common properties between all models
type Model struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type MediaMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AssetPath   string `json:"asset_path"`
}

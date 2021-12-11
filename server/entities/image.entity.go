package entities

import (
	"admire-avatar/config"
)

type BaseImage struct {
	Source string `json:"source"`
}

type Image struct {
	config.Model
	BaseImage
	UserID   uint   `json:"user_id"`
	FolderId uint   `json:"folder_id" bson:"default:0"`
	Type     string `json:"type"`
	Main     bool   `json:"main"`
}

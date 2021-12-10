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
	UserId uint   `json:"user_id"`
	Type   string `json:"type"`
	Main   bool   `json:"main"`
}

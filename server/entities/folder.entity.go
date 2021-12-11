package entities

import (
	"admire-avatar/config"
)

type Folder struct {
	config.Model
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

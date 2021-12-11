package ws

import (
	"admire-avatar/entities"
	"github.com/gorilla/websocket"
)

type Client struct {
	User    *entities.BaseUser
	Conn    *websocket.Conn
	Pool    *Pool
	Channel uint
}

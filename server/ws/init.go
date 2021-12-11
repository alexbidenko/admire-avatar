package ws

import (
	"admire-avatar/entities"
	"admire-avatar/middlewares"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Message struct {
	Body interface{}        `json:"body"`
	User *entities.BaseUser `json:"-"`
}

type MessageBody struct {
	Type    int         `json:"type"`
	Content interface{} `json:"content"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *middlewares.AuthorizedRequest) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func ServeWs(pool *Pool, w http.ResponseWriter, r *middlewares.AuthorizedRequest) {
	conn, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &Client{
		User:    &r.User,
		Conn:    conn,
		Pool:    pool,
		Channel: r.User.ID,
	}

	pool.Register <- client
}

var PrintsPool = NewPool()
var NotificationsPool = NewPool()

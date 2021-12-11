package ws

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Channels   map[uint]map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Channels:   make(map[uint]map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			channel := pool.Channels[client.Channel]
			if channel == nil {
				channel = make(map[*Client]bool)
				pool.Channels[client.Channel] = channel
			}
			channel[client] = true
			break
		case client := <-pool.Unregister:
			channel := pool.Channels[client.Channel]
			if channel != nil && channel[client] {
				delete(channel, client)
			}
			break
		case message := <-pool.Broadcast:
			channel := pool.Channels[message.User.ID]
			if channel != nil {
				for client := range channel {
					if err := client.Conn.WriteJSON(message.Body); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
}

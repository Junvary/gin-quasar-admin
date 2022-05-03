package model

import "sync"

type WsMessage struct {
	Name              string   `json:"name"`
	Avatar            string   `json:"avatar"`
	Text              string   `json:"text"`
	Stamp             string   `json:"stamp"`
	MessageType       string   `json:"message_type"`
	MessageToUserType string   `json:"message_to_user_type"`
	ToUser            []string `json:"toUser"`
}

var (
	//Clients      = make(map[string]*websocket.Conn)
	Clients      sync.Map
	BroadcastMsg = make(chan []byte, 100)
)

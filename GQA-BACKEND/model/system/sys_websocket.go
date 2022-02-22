package system

import "sync"

type WsMessage struct { //接收与返回的结构体
	Name              string   `json:"name"`
	Avatar            string   `json:"avatar"`
	Text              string   `json:"text"`
	Stamp             string   `json:"stamp"`
	MessageType       string   `json:"messageType"`
	MessageToUserType string   `json:"messageToUserType"`
	ToUser            []string `json:"toUser"`
}

var (
	//Clients      = make(map[string]*websocket.Conn)
	Clients      sync.Map
	BroadcastMsg = make(chan []byte, 100)
)

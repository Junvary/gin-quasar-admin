package public

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type ApiWebSocket struct{}

type BackMessage struct { //接收与返回的结构体
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Text   string `json:"text"`
	Stamp  string `json:"stamp"`
}

var (
	clients      = make(map[string]*websocket.Conn)
	broadcastMsg = make(chan []byte, 100)
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *ApiWebSocket) WebSocket(c *gin.Context) {
	go broadcast()

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	
	clientId := uuid.New().String()

	defer func() {
		//连接断开，删除无效client
		delete(clients, clientId)
		ws.Close()
	}()

	clients[clientId] = ws

	for {
		//读取websocket发来的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		var bm BackMessage
		if err = json.Unmarshal(message, &bm); err != nil {
			return
		}
		bm.Stamp = time.Now().Format("2006-01-02 15:04:05")
		bmByte, _ := json.Marshal(bm)
		broadcastMsg <- bmByte
	}
}

func broadcast() {
	for {
		v, ok := <-broadcastMsg
		if !ok {
			break
		}
		go func() {
			for id, client := range clients {
				//tip：id可以判断不给谁发消息
				if err := client.WriteMessage(websocket.TextMessage, v); err != nil {
					//发送失败，客户端异常（断线...）
					delete(clients, id)
				}
			}
		}()
	}
}



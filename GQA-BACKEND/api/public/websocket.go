package public

import (
	"encoding/json"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

type ApiWebSocket struct{}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *ApiWebSocket) WebSocket(c *gin.Context) {
	go servicePublic.ServiceWebSocket.Broadcast()

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// Frontend transmits the username, takes it out, splices the timestamp, and adds _ to prevent duplication
	clientId := c.Param("username") + "_" + strconv.FormatInt(time.Now().Unix(), 10)

	defer func() {
		// Disconnected, delete invalid client
		model.Clients.Delete(clientId)
		_ = ws.Close()
	}()
	model.Clients.Store(clientId, ws)

	for {
		// Read the data sent by websocket
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		var wsMessage model.WsMessage
		if err = json.Unmarshal(message, &wsMessage); err != nil {
			return
		}
		wsMessage.Stamp = time.Now().Format(time.DateTime)
		byteMessage, _ := json.Marshal(wsMessage)
		if wsMessage.MessageType == "chat" {
			model.BroadcastMsg <- byteMessage
		}
	}
}

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
	//前端传递username，取出来和时间戳拼接，加上_是为了放重
	clientId := c.Param("username") + "_" + strconv.FormatInt(time.Now().Unix(), 10)

	defer func() {
		//连接断开，删除无效client
		//delete(system.Clients, clientId)
		model.Clients.Delete(clientId)
		_ = ws.Close()
	}()
	model.Clients.Store(clientId, ws)
	//system.Clients[clientId] = ws

	for {
		//读取websocket发来的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		var wsMessage model.WsMessage
		if err = json.Unmarshal(message, &wsMessage); err != nil {
			return
		}
		wsMessage.Stamp = time.Now().Format("2006-01-02 15:04:05")
		byteMessage, _ := json.Marshal(wsMessage)
		if wsMessage.MessageType == "chat" {
			model.BroadcastMsg <- byteMessage
		}
	}
}

package public

import (
	"encoding/json"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/gorilla/websocket"
	"strings"
)

func Broadcast() {
	for {
		v, ok := <-system.BroadcastMsg
		if !ok {
			break
		}
		go func() {
			var wsMsg system.WsMessage
			_ = json.Unmarshal(v, &wsMsg)
			if wsMsg.MessageType == "chat" {
				//聊天信息群发
				system.Clients.Range(func(id, client interface{}) bool {
					if err := client.(*websocket.Conn).WriteMessage(websocket.TextMessage, v); err != nil {
						system.Clients.Delete(id)
					}
					return true
				})
				//for id, client := range system.Clients {
				//	if err := client.WriteMessage(websocket.TextMessage, v); err != nil {
				//		delete(system.Clients, id)
				//	}
				//}
			} else if wsMsg.MessageType == "notice" {
				if wsMsg.MessageToUserType == "all" {
					//非聊天信息，MessageToUserType是all，那么也是群发
					system.Clients.Range(func(id, client interface{}) bool {
						if err := client.(*websocket.Conn).WriteMessage(websocket.TextMessage, v); err != nil {
							//delete(system.Clients, id)
							system.Clients.Delete(id)
						}
						return true
					})
					//for id, client := range system.Clients {
					//	if err := client.WriteMessage(websocket.TextMessage, v); err != nil {
					//		delete(system.Clients, id)
					//	}
					//}
				} else {
					//非聊天信息，MessageToUserType是some，有ToUser，那么发送给他们
					clientUser := make(map[string]*websocket.Conn)
					for _, u := range wsMsg.ToUser {
						system.Clients.Range(func(id, client interface{}) bool {
							if strings.Contains(id.(string), u+"_") {
								clientUser[id.(string)] = client.(*websocket.Conn)
							}
							return true
						})
						//for id, client := range system.Clients {
						//	if strings.Contains(id, u+"_") {
						//		clientUser[id] = client
						//	}
						//}
					}
					for id, client := range clientUser {
						if err := client.WriteMessage(websocket.TextMessage, v); err != nil {
							//delete(system.Clients, id)
							system.Clients.Delete(id)
						}
					}
				}
			}

		}()
	}
}

package public

import (
	"encoding/json"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gorilla/websocket"
	"strings"
)

type ServiceWebSocket struct{}

func (s *ServiceWebSocket) Broadcast() {
	for {
		v, ok := <-model.BroadcastMsg
		if !ok {
			break
		}
		go func() {
			var wsMsg model.WsMessage
			_ = json.Unmarshal(v, &wsMsg)
			if wsMsg.MessageType == "chat" {
				//聊天信息群发
				model.Clients.Range(func(id, client interface{}) bool {
					if err := client.(*websocket.Conn).WriteMessage(websocket.TextMessage, v); err != nil {
						model.Clients.Delete(id)
					}
					return true
				})
			} else /* if wsMsg.MessageType == "noticeType_message" */ {
				if wsMsg.MessageToUserType == "all" {
					//非聊天信息，MessageToUserType是all，那么也是群发
					model.Clients.Range(func(id, client interface{}) bool {
						if err := client.(*websocket.Conn).WriteMessage(websocket.TextMessage, v); err != nil {
							//delete(system.Clients, id)
							model.Clients.Delete(id)
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
						model.Clients.Range(func(id, client interface{}) bool {
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
							model.Clients.Delete(id)
						}
					}
				}
			}
		}()
	}
}

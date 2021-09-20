package websocketlib

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"goserver/pkg/apiobjects"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
)

type WSMessageWrapper struct {
	Service     string `json:"service"`
	MessageType string `json:"message_type"`
}

// Helper to make Gorilla Websockets threadsafe
type threadSafeWriter struct {
	*websocket.Conn
	sync.Mutex
}

func (t *threadSafeWriter) WriteJSON(v interface{}) error {
	t.Lock()
	defer t.Unlock()
	return t.Conn.WriteJSON(v)
}

func loadServiceAndType(s []byte) (string, string) {
	var wrapper WSMessageWrapper
	err := json.Unmarshal(s, &wrapper)
	if err != nil {
		return "", ""
	}
	return wrapper.Service, wrapper.MessageType
}

// Handle incoming websockets
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request to Websocket
	unsafeConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	c := &threadSafeWriter{unsafeConn, sync.Mutex{}}
	// When this frame returns close the Websocket
	defer c.Close()

	for {
		_, raw, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		_, msgType := loadServiceAndType(raw)
		switch msgType {
		case "LastSeen":
			lastSeen := apiobjects.LastSeen{}
			if err := json.Unmarshal(raw, &lastSeen); err != nil {
				log.Println(err)
				return
			}
			log.Printf("Received Last Seen from user %s \n", lastSeen.UserId)
			message := &apiobjects.LastSeen{
				Service:     "chat",
				MessageType: "LastSeen",
				UserId:      lastSeen.UserId,
				LastTs:      lastSeen.LastTs,
			}
			c.WriteJSON(message)
			break
		case "DeleteMessage":
			delMsg := apiobjects.DeleteMessage{}
			if err := json.Unmarshal(raw, &delMsg); err != nil {
				log.Println(err)
				return
			}
			log.Printf("Received Delete Message for Id %s \n", delMsg.MessageId)
			message := apiobjects.DeleteMessage{
				Service:        "chat",
				MessageType:    "DeleteMessage",
				ConversationId: delMsg.ConversationId,
				MessageId:      delMsg.MessageId,
			}
			c.WriteJSON(message)
			break
		default:
			log.Printf("Unknown message type")
		}
	}
}

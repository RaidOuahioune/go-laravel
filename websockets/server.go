package websockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	upgrader websocket.Upgrader
}

func (m *WebSocketServer) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	m.upgrader = websocket.Upgrader{}
	ws, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

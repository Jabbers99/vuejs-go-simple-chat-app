package handler

import (
	"github.com/gorilla/websocket"
	"chat-app-api/model"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
func HandleWSConn(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	client := &model.Client{Conn: ws, Hub: hub, Send: make(chan []byte, 256)}
	client.Hub.Register <- client
	

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

}
package model

import (
	"time"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Message struct {
	Username string `json:"username"`
	Message string `json:"message"`
	MessageType string `json:"messageType"`
}

func CreateInfoMessage(message string) Message {
	return Message{Message:message, MessageType:"info"}
}
func CreateDefaultMessage(message string, username string) Message {
	return Message{Username: username, Message:message, MessageType:"default"}
}

var Clients = make(map[*websocket.Conn]bool)
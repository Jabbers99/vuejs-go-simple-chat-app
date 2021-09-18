package model
import (
	"bytes"
	"log"
	"time"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {

	Conn *websocket.Conn

	Send chan []byte

	Hub *Hub

	Username string
	Authed bool
}


// WritePump pumps messages from the hub to the websocket Connection.
//
// A goroutine running WritePump is started for each Connection. The
// application ensures that there is at most one writer to a Connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:

			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}



// ReadPump pumps messages from the websocket Connection to the hub.
//
// The application runs ReadPump in a per-Connection goroutine. The application
// ensures that there is at most one reader on a Connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()
	
	//c.Conn.SetReadLimit(maxMessageSize)
	//c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		log.Println("3", err)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		if (c.Authed) {
			msgObj := CreateDefaultMessage(string(message), c.Username)

			msg, _ := json.Marshal(msgObj)
			c.Hub.broadcast <- msg
		} else {

			c.Username = string(message)
			c.Authed = true
			c.Hub.BroadcastInfoMessage(c.Username + " has joined the chat.")
		}
	}
}
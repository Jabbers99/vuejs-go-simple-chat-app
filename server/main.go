package main

import (
	
	"chat-app-api/handler"
	"chat-app-api/model"

	"github.com/gorilla/mux"
	"net/http"

)

func main() {

	hub := model.NewHub()
	go hub.Run()

	router := mux.NewRouter()

	router.HandleFunc("/api/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleWSConn(hub, w, r)
	})
	http.ListenAndServe(":4500", router)
}

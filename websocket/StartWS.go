package ws

import (
	"net/http"
)

func StartWS() {
	wsServer := NewWebsocketServer()
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		(*&w).Header().Set("Access-Control-Allow-Origin", "*")
		(*&w).Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, authorization, x-nextoken")
		(*&w).Header().Set("Access-Control-Allow-Credentials", "true")
		(*&w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		(*&w).Header().Set("Access-Control-Max-Age", "1209600")
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	go http.ListenAndServe(":7002", nil)
}

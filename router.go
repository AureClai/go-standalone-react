package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rakyll/statik/fs"
)

type msg struct {
	Num int
}

func newRouter() *mux.Router {
	// Create the new router
	r := mux.NewRouter()

	// Statik files management
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	// Create the file server from statik
	h := http.FileServer(statikFS)

	r.HandleFunc("/ws", wsHandler)
	r.PathPrefix("/").Handler(h).Methods("GET")

	return r
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go messangeHandling(conn)
}

func messangeHandling(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)
		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
}

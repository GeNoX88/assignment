// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var t = template.Must(template.ParseFiles("./index.html"))

func WS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
}

func home(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", home)
	http.HandleFunc("/ws", WS)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

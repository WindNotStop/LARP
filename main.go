// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var addr = flag.String("addr", "192.168.0.104:8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "static/home.html")
}


func main() {
	flag.Parse()
	roles := []string{"天泽","骚烨","澍妞","韩管家","薛硕"}
	hub := newHub(roles)
	go hub.run()

	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))

	http.Handle("/", http.HandlerFunc(serveHome))
	http.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.URL.Query().Get("role")
		role, _ = url.QueryUnescape(role)
		fmt.Println(role)
		serveWs(role, hub, w, r)
	}))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package LARP

import (
	"flag"
	"github.com/WindNotStop/LARP/model"
	"github.com/go-redis/redis/v7"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var addr = flag.String("addr", "192.168.0.104:8080", "局域网地址")
var scriptName = flag.String("sname", "剧本1", "剧本名")
var client *redis.Client
var script *model.Script

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
	t, err := template.ParseFiles("static/home.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.Execute(w, script)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func mainPage(w http.ResponseWriter, r *http.Request){
	role := r.URL.Query().Get("role")
	role, _ = url.QueryUnescape(role)
	t, err := template.ParseFiles("static/main_page.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.Execute(w, script)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func characterScript(w http.ResponseWriter, r *http.Request){
	role := r.URL.Query().Get("role")
	role, _ = url.QueryUnescape(role)
	t, err := template.ParseFiles("static/character_script.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.Execute(w, script.Characters[role])
	if err != nil {
		log.Fatal(err.Error())
	}
}

func task(w http.ResponseWriter, r *http.Request){
	role := r.URL.Query().Get("role")
	role, _ = url.QueryUnescape(role)
	t, err := template.ParseFiles("static/task.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.Execute(w, script.Characters[role])
	if err != nil {
		log.Fatal(err.Error())
	}
}

func scene(w http.ResponseWriter, r *http.Request){
	arg := r.URL.Query().Get("arg")
	arg, _ = url.QueryUnescape(arg)
	t, err := template.ParseFiles("static/scene.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.Execute(w, script.Scenes[arg])
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	flag.Parse()

	script = &model.Script{}
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := client.Get(*scriptName).Bytes()
	if err != nil {
		log.Fatal(err)
	}
	script.Decode(val)
	hub := newHub(script.Characters)
	go hub.run()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.HandlerFunc(serveHome))
	http.Handle("/characterScript", http.HandlerFunc(characterScript))
	http.Handle("/task", http.HandlerFunc(task))
	http.Handle("/mainPage", http.HandlerFunc(mainPage))
	http.Handle("/scene", http.HandlerFunc(scene))
	http.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.URL.Query().Get("role")
		role, _ = url.QueryUnescape(role)
		serveWs(role, hub, w, r)
	}))
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

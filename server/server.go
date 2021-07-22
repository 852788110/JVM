package server

import (
	"jvmgo/jvm/server/handler"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/upload", handler.ParseCode)
	http.HandleFunc("/index", handler.SayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listen And Server: ", err)
	}
}

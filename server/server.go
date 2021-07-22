package server

import (
	"jvmgo/jvm/server/handler"
	"net/http"
)

func run() {
	http.HandleFunc("/file", handler.ParseCode)
}

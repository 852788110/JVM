package handler

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Good Lucky"))
}

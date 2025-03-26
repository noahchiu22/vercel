package handler

import (
	"net/http"
	"vercel/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := router.Setup_Router()

	server.ServeHTTP(w, r)
}

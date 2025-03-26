package main

import (
	"net/http"
	"vercel/router"
)

func Server(w http.ResponseWriter, r *http.Request) {
	server := router.Setup_Router()

	server.ServeHTTP(w, r)
}

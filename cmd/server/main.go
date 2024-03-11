package main

import (
	"log"
	"proglog/internal/server"
)

var port = ":8080"

func main() {
	srv := server.NewHTTPServer(port)
	log.Fatal(srv.ListenAndServe())
}

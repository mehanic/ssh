package main

import (
	"log"
	"ssh-server/internal/server"
)

func main() {
	srv := server.New("0.0.0.0:2200", "/../../ssh/ssh-server/id_rsa")
	log.Fatal(srv.ListenAndServe())
}

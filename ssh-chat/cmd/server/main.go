package main

import (
	"log"

	"ssh/internal/chat"
	"ssh/internal/territory"

	"github.com/gliderlabs/ssh"
)

func main() {
	availableRooms := []*territory.Room{
		{Name: "a"},
		{Name: "b"},
		{Name: "c"},
	}
	sessions := chat.NewSessionManager(availableRooms)

	ssh.Handle(func(s ssh.Session) {
		chat.ChatHandler(s, sessions)
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}

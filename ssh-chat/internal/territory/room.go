package territory

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type Room struct {
	Name    string
	History []Message
	Users   []User
}

func (r *Room) Enter(sess ssh.Session, term *term.Terminal) {
	u := User{Session: sess, Terminal: term}
	r.Users = append(r.Users, u)
	entryMsg := Message{From: r.Name, Message: "Welcome to my room!"}
	send(u, entryMsg)
	for _, m := range r.History {
		send(u, m)
	}
}

func (r *Room) Leave(sess ssh.Session) {
	r.Users = removeByUsername(r.Users, sess.User())
}

func (r *Room) SendMessage(from, message string) {
	messageObj := Message{From: from, Message: message}
	r.History = append(r.History, messageObj)
	for _, u := range r.Users {
		if u.Session.User() != from {
			send(u, messageObj)
		}
	}
}

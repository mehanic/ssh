package chat

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func ChatHandler(s ssh.Session, sm *SessionManager) {
	term := term.NewTerminal(s, fmt.Sprintf("%s > ", s.User()))
	for {
		line, err := term.ReadLine()
		if err != nil {
			break
		}
		HandleCommand(line, s, term, sm)
	}
}

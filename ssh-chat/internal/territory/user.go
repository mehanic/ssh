package territory

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type User struct {
	Session  ssh.Session
	Terminal *term.Terminal
}

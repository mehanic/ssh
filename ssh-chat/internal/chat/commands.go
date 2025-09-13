package chat

import (
	"regexp"
	"strings"

	"ssh/internal/territory"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

var (
	enterCmd = regexp.MustCompile(`^/enter.*`)
	helpCmd  = regexp.MustCompile(`^/help.*`)
	exitCmd  = regexp.MustCompile(`^/exit.*`)
	listCmd  = regexp.MustCompile(`^/list.*`)
)

func HandleCommand(line string, s ssh.Session, term *term.Terminal, sm *SessionManager) {
	if len(line) == 0 {
		return
	}

	if string(line[0]) != "/" {
		if sm.Sessions[s] != nil {
			sm.Sessions[s].SendMessage(s.User(), line)
		} else {
			term.Write([]byte(helpMsg()))
		}
		return
	}

	switch {
	case exitCmd.MatchString(line):
		return
	case listCmd.MatchString(line):
		term.Write([]byte(listRooms(sm.AvailableRooms)))
	case enterCmd.MatchString(line):
		handleEnter(line, s, term, sm)
	case helpCmd.MatchString(line):
		term.Write([]byte(helpMsg()))
	default:
		term.Write([]byte(helpMsg()))
	}
}

func handleEnter(line string, s ssh.Session, term *term.Terminal, sm *SessionManager) {
	toEnter := strings.Split(line, " ")[1]
	matching := filter(sm.AvailableRooms, func(r *territory.Room) bool {
		return toEnter == r.Name
	})
	if len(matching) == 0 {
		term.Write([]byte("Invalid Room!\n"))
	} else {
		if sm.Sessions[s] != nil {
			sm.Sessions[s].Leave(s)
		}
		r := matching[0]
		r.Enter(s, term)
		sm.Sessions[s] = r
	}
}

func helpMsg() string {
	return `
Hello and welcome to the chat server! Please use
one of the following commands:
	1. /list: To list available rooms
	2. /enter <room>: To enter a room
	3. /exit: To leave the server
	4. /help: To display this message
`
}

func listRooms(rooms []*territory.Room) string {
	var sb strings.Builder
	for _, r := range rooms {
		sb.WriteString(r.Name + "\n")
	}
	return sb.String()
}

func filter[T any](s []T, cond func(t T) bool) []T {
	res := []T{}
	for _, v := range s {
		if cond(v) {
			res = append(res, v)
		}
	}
	return res
}

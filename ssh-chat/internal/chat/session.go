package chat

import (
	"github.com/gliderlabs/ssh"

	"ssh/internal/territory"
)

type SessionManager struct {
	Sessions       map[ssh.Session]*territory.Room
	AvailableRooms []*territory.Room
}

func NewSessionManager(roomList []*territory.Room) *SessionManager {
	return &SessionManager{
		Sessions:       make(map[ssh.Session]*territory.Room),
		AvailableRooms: roomList,
	}
}

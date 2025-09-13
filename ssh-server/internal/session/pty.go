package session

import (
	"encoding/binary"
	"log"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/crypto/ssh"
)

// parseDims extracts terminal dimensions (width x height) from the provided buffer.
func parseDims(b []byte) (uint32, uint32) {
	w := binary.BigEndian.Uint32(b)
	h := binary.BigEndian.Uint32(b[4:])
	return w, h
}

// Winsize stores the Height and Width of a terminal.
type Winsize struct {
	Height uint16
	Width  uint16
	x      uint16
	y      uint16
}

// SetWinsize sets the size of the given pty.
func SetWinsize(fd uintptr, w, h uint32) {
	ws := &Winsize{Width: uint16(w), Height: uint16(h)}
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		fd,
		uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(ws)))
	if errno != 0 {
		log.Printf("Error resizing pty: %v", errno)
	}
}

func handleRequests(reqs <-chan *ssh.Request, bashf *os.File) {
	for req := range reqs {
		switch req.Type {
		case "shell":
			if len(req.Payload) == 0 {
				req.Reply(true, nil)
			}
		case "pty-req":
			termLen := req.Payload[3]
			w, h := parseDims(req.Payload[termLen+4:])
			SetWinsize(bashf.Fd(), w, h)
			req.Reply(true, nil)
		case "window-change":
			w, h := parseDims(req.Payload)
			SetWinsize(bashf.Fd(), w, h)
		}
	}
}

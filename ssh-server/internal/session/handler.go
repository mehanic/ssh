package session

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
	"syscall"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh"
)

func HandleChannels(chans <-chan ssh.NewChannel) {
	for newChannel := range chans {
		go handleChannel(newChannel)
	}
}

func handleChannel(newChannel ssh.NewChannel) {
	if t := newChannel.ChannelType(); t != "session" {
		newChannel.Reject(ssh.UnknownChannelType, fmt.Sprintf("unknown channel type: %s", t))
		return
	}

	connection, requests, err := newChannel.Accept()
	if err != nil {
		log.Printf("Could not accept channel (%s)", err)
		return
	}

	bash := exec.Command("bash")
	bash.SysProcAttr = &syscall.SysProcAttr{
		Setctty: true,
		Setsid:  true,
	}

	close := func() {
		connection.Close()
		if bash.Process != nil {
			_, err := bash.Process.Wait()
			if err != nil {
				log.Printf("Failed to exit bash (%s)", err)
			}
		}
		log.Printf("Session closed")
	}

	log.Print("Creating pty...")
	bashf, err := pty.Start(bash)
	if err != nil {
		log.Printf("Could not start pty (%s)", err)
		close()
		return
	}

	var once sync.Once
	go func() {
		io.Copy(connection, bashf)
		once.Do(close)
	}()
	go func() {
		io.Copy(bashf, connection)
		once.Do(close)
	}()

	go handleRequests(requests, bashf)
}

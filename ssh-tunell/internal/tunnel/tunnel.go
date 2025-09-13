package tunnel

import (
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func Start(client *ssh.Client, local, remote string) error {
	pipe := func(writer, reader net.Conn) {
		defer writer.Close()
		defer reader.Close()
		if _, err := io.Copy(writer, reader); err != nil {
			log.Printf("failed to copy: %s", err)
		}
	}

	listener, err := net.Listen("tcp", local)
	if err != nil {
		return err
	}
	for {
		here, err := listener.Accept()
		if err != nil {
			return err
		}
		go func(here net.Conn) {
			there, err := client.Dial("tcp", remote)
			if err != nil {
				log.Printf("failed to dial to remote: %q", err)
				here.Close()
				return
			}
			go pipe(there, here)
			go pipe(here, there)
		}(here)
	}
}

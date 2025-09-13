package connection

import (
	"golang.org/x/crypto/ssh"
	"log"
)

func Connect(addr string, config *ssh.ClientConfig) *ssh.Client {
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	return client
}

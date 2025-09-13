package sshconn

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func Connect(host, port, user string, signer ssh.Signer) *ssh.Client {
	clientConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), clientConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	return connection
}

func RunCommand(client *ssh.Client, command string) string {
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	return string(output)
}

package sshconn

import (
	"io"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func CreateSSHClientConfig(user string, key ssh.Signer) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
}

func ConnectSSH(host, port string, config *ssh.ClientConfig) *ssh.Client {
	address := net.JoinHostPort(host, port)
	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %s", host, err)
	}
	return client
}

func RunRemoteCommand(client *ssh.Client, command string) string {
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to get stdout: %s", err)
	}

	if err := session.Start(command); err != nil {
		log.Fatalf("Failed to start command: %s", err)
	}

	output, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatalf("Failed to read stdout: %s", err)
	}

	if err := session.Wait(); err != nil {
		log.Fatalf("Command failed: %s", err)
	}

	return string(output)
}

package auth

import (
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

const EnvSSHAuthSock = "SSH_AUTH_SOCK"

func Agent() (ssh.AuthMethod, error) {
	conn, err := net.Dial("unix", os.Getenv(EnvSSHAuthSock))
	if err != nil {
		return nil, err
	}
	client := agent.NewClient(conn)
	return ssh.PublicKeysCallback(client.Signers), nil
}

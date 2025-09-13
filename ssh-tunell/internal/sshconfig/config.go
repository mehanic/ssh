package sshconfig

import (
	"os/user"

	"golang.org/x/crypto/ssh"
)

func New(methods ...ssh.AuthMethod) (*ssh.ClientConfig, error) {
	current, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &ssh.ClientConfig{
		User:            current.Username,
		Auth:            methods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}

func NewWithUser(username string, methods ...ssh.AuthMethod) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:            username,
		Auth:            methods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

package sshconfig

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

// PrivateKeyFile повертає ssh.AuthMethod для приватного ключа
func PrivateKeyFile(file string) ssh.AuthMethod {
	buffer, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read private key: %s", err)
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		log.Fatalf("Unable to parse private key: %s", err)
	}
	return ssh.PublicKeys(key)
}

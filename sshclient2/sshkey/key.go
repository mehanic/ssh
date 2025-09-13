package sshkey

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func LoadPrivateKey(path string) ssh.Signer {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		log.Fatalf("Unable to parse private key: %v", err)
	}

	return signer
}

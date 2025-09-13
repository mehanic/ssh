package sshkey

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func LoadPrivateKey(path string, passphrase []byte) ssh.Signer {
	keyBytes, err := ioutil.ReadFile(os.ExpandEnv(path))
	if err != nil {
		log.Fatalf("Failed to read private key: %s", err)
	}

	var key ssh.Signer
	if len(passphrase) > 0 {
		key, err = ssh.ParsePrivateKeyWithPassphrase(keyBytes, passphrase)
	} else {
		key, err = ssh.ParsePrivateKey(keyBytes)
	}
	if err != nil {
		log.Fatalf("Failed to parse private key: %s", err)
	}

	return key
}

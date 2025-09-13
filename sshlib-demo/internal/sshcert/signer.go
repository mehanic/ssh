package sshcert

import (
	"crypto/rsa"
	"fmt"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

const formatCert = "ssh-rsa-cert-v01@openssh.com"

func GenerateSignerFromKey(priv *rsa.PrivateKey) (ssh.Signer, error) {
	return ssh.NewSignerFromKey(priv)
}

func GenerateSignerFromBytes(data []byte) (ssh.Signer, error) {
	return ssh.ParsePrivateKey(data)
}

func ListAndCast(keys []*agent.Key) error {
	for _, key := range keys {
		pub, err := ssh.ParsePublicKey(key.Blob)
		if err != nil {
			return err
		}
		if key.Format == formatCert {
			cert, ok := pub.(*ssh.Certificate)
			if !ok {
				return fmt.Errorf("failed to cast key to certificate")
			}
			fmt.Println("Found certificate:", cert.Key)
		}
	}
	return nil
}

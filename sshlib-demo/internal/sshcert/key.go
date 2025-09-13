package sshcert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

const keySize = 2048
const typePrivateKey = "RSA PRIVATE KEY"

func GenerateKey() (*rsa.PrivateKey, ssh.PublicKey, error) {
	priv, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, err
	}
	pub, err := ssh.NewPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return priv, pub, nil
}

func MarshalRSAPrivate(priv *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  typePrivateKey,
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})
}

func UnmarshalRSAPrivate(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func MarshalRSAPublic(pub ssh.PublicKey) []byte {
	return bytes.TrimSuffix(ssh.MarshalAuthorizedKey(pub), []byte{'\n'})
}

func UnmarshalRSAPublic(data []byte) (ssh.PublicKey, error) {
	pub, _, _, _, err := ssh.ParseAuthorizedKey(data)
	return pub, err
}

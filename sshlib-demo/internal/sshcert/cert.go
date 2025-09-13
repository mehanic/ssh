package sshcert

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func GenerateCert(pub ssh.PublicKey) *ssh.Certificate {
	return &ssh.Certificate{
		CertType: ssh.UserCert,
		Key:      pub,
		Permissions: ssh.Permissions{
			CriticalOptions: map[string]string{},
			Extensions:      map[string]string{"permit-agent-forwarding": ""},
		},
	}
}

func MarshalCert(cert *ssh.Certificate) []byte {
	return ssh.MarshalAuthorizedKey(cert)
}

func UnmarshalCert(data []byte) (*ssh.Certificate, error) {
	pub, _, _, _, err := ssh.ParseAuthorizedKey(data)
	if err != nil {
		return nil, err
	}
	cert, ok := pub.(*ssh.Certificate)
	if !ok {
		return nil, fmt.Errorf("failed to cast to certificate")
	}
	return cert, nil
}

func GenerateAndSign() (*rsa.PrivateKey, *ssh.Certificate, error) {
	priv, pub, err := GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	signer, err := GenerateSignerFromKey(priv)
	if err != nil {
		return nil, nil, err
	}
	cert := GenerateCert(pub)
	return priv, cert, cert.SignCert(rand.Reader, signer)
}

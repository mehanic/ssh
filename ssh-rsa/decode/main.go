package decode

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func EncryptRSAOAEP(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, data, nil)
}

package download

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func DecodeRSAKey(bytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func main() {
	keyBytes, _ := ioutil.ReadFile("rsa.priv")
	privKey, _ := DecodeRSAKey(keyBytes)
	fmt.Println(privKey) // пример, чтобы проверить что ключ загружен
}

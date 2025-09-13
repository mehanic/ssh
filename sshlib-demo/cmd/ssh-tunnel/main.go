package main

import (
	"fmt"
	"log"

	"sshlib-demo/internal/sshcert"
	//"golang.org/x/crypto/ssh"
)

func main() {
	priv, cert, err := sshcert.GenerateAndSign()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private Key:\n%s\n", sshcert.MarshalRSAPrivate(priv))
	fmt.Printf("Cert:\n%s\n", sshcert.MarshalCert(cert))

	// signer, err := sshcert.GenerateSignerFromKey(priv)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// certSigner, err := ssh.NewCertSigner(cert, signer)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cfg := &ssh.ClientConfig{
	// 	User:            "root",
	// 	Auth:            []ssh.AuthMethod{ssh.PublicKeys(certSigner)},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	// serverAddr := "127.0.0.1:22"
	// client, err := ssh.Dial("tcp", serverAddr, cfg)
	// if err != nil {
	// 	log.Fatal("Failed to connect to SSH server:", err)
	// }
	// defer client.Close()
	// fmt.Println("SSH connection established to", serverAddr)

}

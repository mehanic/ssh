package session

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func StartShell(client *ssh.Client) {
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("linux", 80, 40, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}

	if err := session.Wait(); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
}

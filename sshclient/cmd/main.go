package main

import (
	"sshclient/config"
	"sshclient/connection"
	"sshclient/session"
)

func main() {
	sshConfig := config.NewClientConfig("username", "password")
	client := connection.Connect("192.168.3.111:22", sshConfig)
	session.StartShell(client)
}

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"

// 	"golang.org/x/crypto/ssh"
// 	"golang.org/x/crypto/ssh/knownhosts"
// )

// func main() {
// 	key, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
// 	if err != nil {
// 		panic(fmt.Sprintf("не удалось прочитать ключ: %v", err))
// 	}

// 	signer, err := ssh.ParsePrivateKey(key)
// 	if err != nil {
// 		panic(fmt.Sprintf("не удалось распарсить ключ: %v", err))
// 	}

// 	hostKeyCallback, err := knownhosts.New(os.Getenv("HOME") + "/.ssh/known_hosts")
// 	if err != nil {
// 		panic(fmt.Sprintf("не удалось загрузить known_hosts: %v", err))
// 	}

// 	config := &ssh.ClientConfig{
// 		User: "ubuntu",
// 		Auth: []ssh.AuthMethod{
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: hostKeyCallback,
// 	}

// 	client, err := ssh.Dial("tcp", "example.com:22", config)
// 	if err != nil {
// 		panic(fmt.Sprintf("не удалось подключиться: %v", err))
// 	}
// 	defer client.Close()

// 	session, err := client.NewSession()
// 	if err != nil {
// 		panic(fmt.Sprintf("не удалось создать сессию: %v", err))
// 	}
// 	defer session.Close()

// 	output, err := session.CombinedOutput("ls -la")
// 	if err != nil {
// 		panic(fmt.Sprintf("ошибка при выполнении команды: %v", err))
// 	}

// 	fmt.Println("Результат команды:\n", string(output))
// }

//https://www.pixelstech.net/article/1699714722-guide-to-implement-an-ssh-client-using-golang

package main

import (
	"fmt"
	"sshclient1/sshconfig"
	"sshclient1/sshconn"
	"sshclient1/sshkey"
)

func main() {
	host := "myhost"
	command := "ls -la"

	user, hostName, port := sshconfig.LoadSSHConfig(host)
	key := sshkey.LoadPrivateKey("$HOME/.ssh/id_rsa", nil)
	//key := sshkey.LoadPrivateKey("$HOME/.ssh/id_rsa", []byte("keyOfPassword"))
	clientConfig := sshconn.CreateSSHClientConfig(user, key)

	client := sshconn.ConnectSSH(hostName, port, clientConfig)
	defer client.Close()

	output := sshconn.RunRemoteCommand(client, command)
	fmt.Println(output)
}

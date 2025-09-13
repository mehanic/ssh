package main

import (
	"fmt"

	"sshclient2/sshconfig"
	"sshclient2/sshconn"
	"sshclient2/sshkey"
)

func main() {
	hostAlias := "example"

	cfg := sshconfig.LoadSSHConfig(hostAlias)

	signer := sshkey.LoadPrivateKey(cfg.IdentityFile)

	client := sshconn.Connect(cfg.HostName, cfg.Port, cfg.User, signer)
	defer client.Close()

	output := sshconn.RunCommand(client, "uname -a")
	fmt.Println(output)
}

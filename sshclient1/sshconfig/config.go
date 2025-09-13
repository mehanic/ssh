package sshconfig

import (
	"log"
	"os"

	"github.com/kevinburke/ssh_config"
)

func LoadSSHConfig(host string) (user, hostName, port string) {
	f, err := os.Open(os.ExpandEnv("$HOME/.ssh/config"))
	if err != nil {
		log.Fatalf("Failed to open SSH config: %s", err)
	}
	defer f.Close()

	cfg, err := ssh_config.Decode(f)
	if err != nil {
		log.Fatalf("Failed to decode SSH config: %s", err)
	}

	user, _ = cfg.Get(host, "User")
	hostName, _ = cfg.Get(host, "HostName")
	port, _ = cfg.Get(host, "Port")
	if port == "" {
		port = "22"
	}

	return
}

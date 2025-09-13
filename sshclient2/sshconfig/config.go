package sshconfig

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kevinburke/ssh_config"
)

type SSHConfig struct {
	Host         string
	HostName     string
	User         string
	Port         string
	IdentityFile string
}

func LoadSSHConfig(host string) SSHConfig {
	sshConfigPath := filepath.Join(os.Getenv("HOME"), ".ssh", "config")

	configFile, err := os.Open(sshConfigPath)
	if err != nil {
		log.Fatalf("Failed to open SSH config file: %v", err)
	}
	defer configFile.Close()

	cfg, err := ssh_config.Decode(configFile)
	if err != nil {
		log.Fatalf("Failed to parse SSH config file: %v", err)
	}

	hostname, _ := cfg.Get(host, "HostName")
	if hostname == "" {
		log.Fatalf("HostName not found for host: %s", host)
	}

	user, _ := cfg.Get(host, "User")
	if user == "" {
		user = os.Getenv("USER")
	}

	port, _ := cfg.Get(host, "Port")
	if port == "" {
		port = "22"
	}

	identityFile, _ := cfg.Get(host, "IdentityFile")
	if identityFile == "" {
		identityFile = filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
	}

	return SSHConfig{
		Host:         host,
		HostName:     hostname,
		User:         user,
		Port:         port,
		IdentityFile: identityFile,
	}
}

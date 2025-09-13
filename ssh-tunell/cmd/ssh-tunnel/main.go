// package main

// import (
// 	"log"

// 	"golang.org/x/crypto/ssh"

// 	"ssh-tunell/internal/auth"
// 	"ssh-tunell/internal/sshconfig"
// 	"ssh-tunell/internal/tunnel"
// )

// func main() {
// 	// auth methods
// 	authInteractive := auth.Interactive()
// 	authAgent, err := auth.Agent()
// 	if err != nil {
// 		log.Fatalf("failed to connect to ssh-agent: %v", err)
// 	}

// 	// ssh config
// 	clientConfig, err := sshconfig.New(authAgent, authInteractive)
// 	if err != nil {
// 		log.Fatalf("failed to create ssh config: %v", err)
// 	}

// 	// connect
// 	clientConn, err := ssh.Dial("tcp", "127.0.0.1:22", clientConfig)
// 	if err != nil {
// 		log.Fatalf("failed to connect to ssh server: %v", err)
// 	}

// 	// start tunnel
// 	if err := tunnel.Start(clientConn, "localhost:1600", "localhost:1500"); err != nil {
// 		log.Fatalf("failed to tunnel traffic: %v", err)
// 	}
// }

package main

import (
	"flag"
	"log"

	"golang.org/x/crypto/ssh"

	"ssh-tunell/internal/auth"
	"ssh-tunell/internal/sshconfig"
	"ssh-tunell/internal/tunnel"
)

type Config struct {
	User   string
	Host   string
	Local  string
	Remote string
}

func parseFlags() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.User, "user", "", "SSH username (example: root, ubuntu)")
	flag.StringVar(&cfg.Host, "host", "127.0.0.1:22", "SSH server address (host:port)")
	flag.StringVar(&cfg.Local, "L", "localhost:1600", "Local bind address")
	flag.StringVar(&cfg.Remote, "R", "localhost:1500", "Remote target address")
	flag.Parse()
	return cfg
}

func main() {
	cfg := parseFlags()

	// auth methods
	authInteractive := auth.Interactive()
	authAgent, err := auth.Agent()
	if err != nil {
		log.Fatalf("failed to connect to ssh-agent: %v", err)
	}

	// ssh config
	var clientConfig *ssh.ClientConfig
	if cfg.User != "" {
		clientConfig = sshconfig.NewWithUser(cfg.User, authAgent, authInteractive)
	} else {
		clientConfig, err = sshconfig.New(authAgent, authInteractive)
		if err != nil {
			log.Fatalf("failed to create ssh config: %v", err)
		}
	}

	// connect
	clientConn, err := ssh.Dial("tcp", cfg.Host, clientConfig)
	if err != nil {
		log.Fatalf("failed to connect to ssh server: %v", err)
	}

	// start tunnel
	log.Printf("Starting tunnel: %s -> %s via %s@%s", cfg.Local, cfg.Remote, clientConfig.User, cfg.Host)
	if err := tunnel.Start(clientConn, cfg.Local, cfg.Remote); err != nil {
		log.Fatalf("failed to tunnel traffic: %v", err)
	}
}

//go run main.go -user root -host 192.168.1.10:22 -L localhost:1600 -R localhost:1500

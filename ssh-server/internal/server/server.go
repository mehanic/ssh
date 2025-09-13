package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"ssh-server/internal/session"

	"golang.org/x/crypto/ssh"
)

type Server struct {
	addr   string
	config *ssh.ServerConfig
}

func New(addr, keyPath string) *Server {
	privateBytes, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("Failed to load private key (%s)", keyPath)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}

	config := &ssh.ServerConfig{
		PasswordCallback: passwordAuth,
	}
	config.AddHostKey(private)

	return &Server{addr: addr, config: config}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.addr, err)
	}

	log.Printf("Listening on %s...", s.addr)
	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
			continue
		}

		sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, s.config)
		if err != nil {
			log.Printf("Handshake failed: %s", err)
			continue
		}

		log.Printf("New SSH connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())

		go ssh.DiscardRequests(reqs)
		go session.HandleChannels(chans)
	}
}

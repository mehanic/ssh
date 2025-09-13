package tunnel

import (
	"io"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"

	"ssh-tunnel/endpoint"
)

type SSHTunnel struct {
	Local  *endpoint.Endpoint
	Server *endpoint.Endpoint
	Remote *endpoint.Endpoint
	Config *ssh.ClientConfig
	Log    *log.Logger
}

func New(tunnel string, auth ssh.AuthMethod, destination string) *SSHTunnel {
	server := endpoint.New(tunnel)
	if server.Port == 0 {
		server.Port = 22
	}
	return &SSHTunnel{
		Local:  endpoint.New("localhost:0"),
		Server: server,
		Remote: endpoint.New(destination),
		Config: &ssh.ClientConfig{
			User: server.User,
			Auth: []ssh.AuthMethod{auth},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         5 * time.Second,
		},
	}
}

func (t *SSHTunnel) logf(format string, args ...interface{}) {
	if t.Log != nil {
		t.Log.Printf(format, args...)
	}
}

func (t *SSHTunnel) Start() error {
	listener, err := net.Listen("tcp", t.Local.String())
	if err != nil {
		return err
	}
	defer listener.Close()

	t.Local.Port = listener.Addr().(*net.TCPAddr).Port

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		t.logf("accepted connection")
		go t.forward(conn)
	}
}

func (t *SSHTunnel) forward(localConn net.Conn) {
	serverConn, err := ssh.Dial("tcp", t.Server.String(), t.Config)
	if err != nil {
		t.logf("server dial error: %s", err)
		return
	}
	t.logf("connected to %s (1 of 2)\n", t.Server.String())

	remoteConn, err := serverConn.Dial("tcp", t.Remote.String())
	if err != nil {
		t.logf("remote dial error: %s", err)
		return
	}
	t.logf("connected to %s (2 of 2)\n", t.Remote.String())

	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			t.logf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}

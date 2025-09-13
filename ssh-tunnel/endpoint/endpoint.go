package endpoint

import (
	"fmt"
	"strconv"
	"strings"
)

type Endpoint struct {
	Host string
	Port int
	User string
}

func New(s string) *Endpoint {
	e := &Endpoint{Host: s}

	// user@host
	if parts := strings.Split(e.Host, "@"); len(parts) > 1 {
		e.User = parts[0]
		e.Host = parts[1]
	}

	// host:port
	if parts := strings.Split(e.Host, ":"); len(parts) > 1 {
		e.Host = parts[0]
		e.Port, _ = strconv.Atoi(parts[1])
	}

	return e
}

func (e *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

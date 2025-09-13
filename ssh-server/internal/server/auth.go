package server

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

// func passwordAuth(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
// 	if c.User() == "foo" && string(pass) == "bar" {
// 		return nil, nil
// 	}
// 	return nil, fmt.Errorf("password rejected for %q", c.User())
// }

func passwordAuth(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
    log.Printf("Auth attempt: user=%s password=%q", c.User(), pass)
    if c.User() == "foo" && string(pass) == "bar" {
        return nil, nil
    }
    return nil, fmt.Errorf("password rejected for %q", c.User())
}

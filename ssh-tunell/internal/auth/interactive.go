package auth

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func Interactive() ssh.AuthMethod {
	return ssh.KeyboardInteractive(
		func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			if len(questions) == 0 {
				fmt.Printf("%s %s\n", user, instruction)
			}
			answers := make([]string, len(questions))

			for i, question := range questions {
				fmt.Print(question)
				if echos[i] {
					if _, err := fmt.Scan(&answers[i]); err != nil {
						return nil, err
					}
				} else {
					answer, err := term.ReadPassword(syscall.Stdin)
					if err != nil {
						return nil, err
					}
					answers[i] = string(answer)
				}
			}
			return answers, nil
		})
}

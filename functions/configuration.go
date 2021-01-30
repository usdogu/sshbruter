package functions

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

func New(pwd string, username string, timeout time.Duration) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
		User:            username,
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
	}
}

func Connect(host string, port int, config *ssh.ClientConfig) (bool, error) {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return false, err
	}
	client.Close()
	return true, nil

}

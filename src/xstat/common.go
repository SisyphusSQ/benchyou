/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xstat

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
)

func splitColumns(line string) []string {
	cols := make([]string, 0)
	for _, f := range strings.Split(line, " ") {
		if len(f) > 0 {
			cols = append(cols, f)
		}
	}
	return cols
}

func sshConnect(user, password, host string, port int) (client *ssh.Client, err error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	dsn := fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", dsn, sshConfig); err != nil {
		return
	}

	/*
		if session, err = client.NewSession(); err != nil {
			client.Close()
			return
		}
	*/

	return
}

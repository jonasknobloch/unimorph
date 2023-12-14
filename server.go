package unimorph

import (
	"net"
	"os"
)

func Server(name string, split func(form string) []byte) error {
	err := os.Remove(name)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	l, err := net.Listen("unix", name)

	if err != nil {
		return err
	}

	defer l.Close()

	for {
		c, err := l.Accept()

		if err != nil {
			return err
		}

		defer c.Close()

		buffer := make([]byte, 1024)

		n, err := c.Read(buffer)

		if err != nil {
			return err
		}

		_, err = c.Write(split(string(buffer[:n])))

		if err != nil {
			return err
		}
	}
}

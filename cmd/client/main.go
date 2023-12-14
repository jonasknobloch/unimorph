package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	c, err := net.Dial("unix", "/tmp/unimorph.sock")

	if err != nil {
		log.Fatalln(err)
	}

	defer c.Close()

	m := []byte("afghánský")

	_, err = c.Write(m)

	if err != nil {
		log.Fatalln(err)
	}

	b := make([]byte, 1024)

	n, err := c.Read(b)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b[:n]))
}

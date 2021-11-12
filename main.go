package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	s := flag.String("socket", "/tmp/go-ipc-pipe.sock", "Specifies the socket to connect to")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("<-")
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		c, err := net.Dial("unix", *s)
		if err != nil {
			panic(err)
		}
		_, _ = c.Write([]byte(line))

		b, err := ioutil.ReadAll(c)
		if err != nil {
			panic(err)
		}
		fmt.Println("->" + string(b))
	}
}

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
	s := flag.String("socket", "", "Specifies the socket to connect to")
	flag.Parse()

	if s == nil || *s == "" {
		panic("-socket parameter is required")
	}

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

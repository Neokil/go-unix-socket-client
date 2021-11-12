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
	c := flag.String("command", "", "Specifies a command to run. If not set interactive mode will be used.")
	flag.Parse()

	if s == nil || *s == "" {
		panic("-socket parameter is required")
	}

	if c != nil && *c != "" {
		executeCommand(*c, *s)
		return
	}

	interactiveMode(*s)
}

func executeCommand(cmd string, socket string) {
	c, err := net.Dial("unix", socket)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = c.Write([]byte(cmd + "\n"))
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(c)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(line)
}

func interactiveMode(socket string) {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("<-")
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		c, err := net.Dial("unix", socket)
		if err != nil {
			panic(err)
		}
		defer c.Close()
		_, _ = c.Write([]byte(line))

		b, err := ioutil.ReadAll(c)
		if err != nil {
			panic(err)
		}
		fmt.Println("->" + string(b))
	}
}

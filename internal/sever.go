package webAppCalculator

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	ConnHost string
	ConnPort string
	ConnType string
}

func ErrHandler(err error) {
	log.Fatal(err)
}

func (s *Server) Run() {
	server := Server{
		ConnHost: "localhost",
		ConnPort: "8080",
		ConnType: "tcp",
	}

	listener, err := net.Listen(server.ConnType, server.ConnHost + ":" + server.ConnPort)
	ErrHandler(err)
	defer listener.Close()

	for {
		con, err := listener.Accept()
		ErrHandler(err)
		fmt.Println("client connected")

		client := &Client{
			Con: con,
		}

		go client.HandleRequest()
	}
}

func (c *Client) HandleRequest() {
	reader := bufio.NewReader(c.Con)
	DoMath(c.Exp)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client left")
			c.Con.Close()
			return
		}

		fmt.Printf("massage: %s", string(data))
		c.Con.Write([]byte("massage received \n"))
	}
}

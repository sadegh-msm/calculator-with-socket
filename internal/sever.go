package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"webAppCalculator/internal/calculation"
)

// Server type that carries all info about server stats
type Server struct {
	ConnHost string
	ConnPort string
	ConnType string
}

// ErrHandler this function gives an error and log it where ever its called
func ErrHandler(err error) {
	log.Fatal(err)
}

// RunServer this method starts a new server for clients to connect to it
func (s *Server) RunServer(connType string, connHost string, connPort string) {
	fmt.Println("server is up")
	listener, err := net.Listen(connType, connHost+":"+connPort)
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

// HandleRequest this method handles the request that comes from client
func (c *Client) HandleRequest() {
	reader := bufio.NewReader(c.Con)
	calculation.DoMath(Context)

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

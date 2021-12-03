package webAppCalculator

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	Con net.Conn
	Exp string
}

var Context string

// RunClient make a new client to connect to the server
// client can do simple math calculation
func (s *Server) RunClient(connType string, connHost string, connPort string) {
	con, err := net.Dial(connType, connHost+":"+connPort)
	ErrHandler(err)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("please first enter the number after operator and second number")

		Context, _ = reader.ReadString('\n')

		result, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Println("the result is: " + result)
	}
}

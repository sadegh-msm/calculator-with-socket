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

func (c *Client) ServeHttp() {
	server := Server{
		ConnHost: "localhost",
		ConnPort: "8080",
		ConnType: "tcp",
	}
	con, err := net.Dial(server.ConnType, server.ConnHost + ":" + server.ConnPort)
	 ErrHandler(err)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("please first enter the number after operator and second number")

		c.Exp, _ = reader.ReadString('\n')

		result, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Println("the result is: " + result)
	}
}

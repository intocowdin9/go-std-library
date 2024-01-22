package main

import (
	"flag"
	"fmt"
)

func main() {

	var username *string = flag.String("username", "return", "database username")
	var password *string = flag.String("password", "password", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *string = flag.String("port", "9000", "database port")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)
}

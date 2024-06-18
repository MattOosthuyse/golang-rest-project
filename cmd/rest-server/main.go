package main

import (
	"cmd/rest-server/internal/server"
	"fmt"
)

func main() {
	fmt.Println("Hello from main")
	server.RestServer()
}

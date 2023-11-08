package main

import (
	"os"

	"github.com/Kaspetti/League-Vis/internal/server"
)


func main() {
    ip := os.Args[1]
    port := os.Args[2]
    server.RunServer(ip, port)
}

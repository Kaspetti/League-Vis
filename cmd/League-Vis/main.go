package main

import (
	"log"
	"os"

	"github.com/Kaspetti/League-Vis/internal/server"
)


func main() {
    ip, ok := os.LookupEnv("IP")
    if !ok {
        log.Fatalln("IP not found in environment...")
    }

    port, ok := os.LookupEnv("PORT")
    if !ok {
        log.Fatalln("PORT not found in environment...")
    }

    server.RunServer(ip, port)
}

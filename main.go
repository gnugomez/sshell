package main

import (
	"flag"
	"sshell/server"

	"github.com/charmbracelet/log"
)

func main() {
	config := server.Config{Port: 2222}

	flag.StringVar(&config.HostKeyPath, "key", "", "Path to the host key file")
	flag.IntVar(&config.Port, "port", 2222, "Port number to listen on")
	flag.Parse()

	log.SetLevel(log.DebugLevel)

	server := server.CreateServer(config)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server", "error", err)
	}
}

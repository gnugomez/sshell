package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

func main() {
	config := Config{Port: 2222}

	flag.StringVar(&config.ScriptPath, "script", "", "Path to the shell script file")
	flag.StringVar(&config.HostKeyPath, "key", "", "Path to the host key file")
	flag.IntVar(&config.Port, "port", 2222, "Port number to listen on")
	flag.Parse()

	if config.ScriptPath == "" {
		logrus.Fatal("No script file provided. Use the -script flag to specify the shell script file.")
	}

	server := CreateServer(config)
	if err := server.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
}

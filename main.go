package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sshell/tui"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

func main() {
	var (
		hostKeyPath string
		port        int
	)

	flag.StringVar(&hostKeyPath, "key", "", "Path to the host key file")
	flag.IntVar(&port, "port", 2222, "Port number to listen on")
	flag.Parse()

	log.SetLevel(log.DebugLevel)

	// Create SSH server
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf(":%d", port)),
		wish.WithHostKeyPath(hostKeyPath),
		// Add logging middleware
		wish.WithMiddleware(
			logging.Middleware(),
			bubbletea.Middleware(teaMiddleware),
		),
	)

	if err != nil {
		log.Fatal("Failed to create server", "error", err)
	}

	// Handle graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting SSH server", "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && err != ssh.ErrServerClosed {
			log.Fatal("Failed to start server", "error", err)
		}
	}()

	// Wait for shutdown signal
	<-done
	log.Info("Shutting down server...")

	// Give outstanding connections 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Error("Server shutdown failed", "error", err)
	}
}

func teaMiddleware(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	return tui.CreateMenu(), []tea.ProgramOption{tea.WithAltScreen()}
}

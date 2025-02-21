package server

import (
	"fmt"
	"sshell/menu"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/log"
	"github.com/gliderlabs/ssh"
)

// Config holds server configuration.
type Config struct {
	HostKeyPath string
	Port        int
}

// Server represents our SSH server.
type Server struct {
	config Config
}

// CreateServer returns a new Server instance.
func CreateServer(config Config) *Server {
	return &Server{config: config}
}

// Start launches the SSH server.
func (s *Server) Start() error {
	log.Info("Starting SSH server", "port", s.config.Port)

	var serverOptions []ssh.Option
	if s.config.HostKeyPath != "" {
		serverOptions = append(serverOptions, ssh.HostKeyFile(s.config.HostKeyPath))
	}

	address := fmt.Sprintf(":%d", s.config.Port)
	return ssh.ListenAndServe(address, s.handleSession, serverOptions...)
}

// Update handleSession function
func (s *Server) handleSession(session ssh.Session) {
	log.Info("New SSH session started", "user", session.User())
	defer session.Close()

	if _, _, isPty := session.Pty(); !isPty {
		session.Write([]byte("PTY required for this session\n"))
		session.Exit(1)
		return
	}

	p := tea.NewProgram(
		menu.CreateMenu(),
		tea.WithInput(session),
		tea.WithOutput(session),
	)

	// Handle window size changes
	_, winCh, _ := session.Pty()
	go func() {
		for win := range winCh {
			p.Send(tea.WindowSizeMsg{
				Width:  win.Width,
				Height: win.Height,
			})
		}
	}()

	if _, err := p.Run(); err != nil {
		log.Error("Menu error", "error", err)
	}
	session.Close()
}

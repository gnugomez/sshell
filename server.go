package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ScriptPath  string
	HostKeyPath string
	Port        int
}

type Server struct {
	config Config
	logger *logrus.Logger
}

func CreateServer(config Config) *Server {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	return &Server{
		config: config,
		logger: logger,
	}
}

func (s *Server) Start() error {
	s.logger.Info("Starting SSH server on port ", s.config.Port)

	var serverOptions []ssh.Option
	if s.config.HostKeyPath != "" {
		serverOptions = append(serverOptions, ssh.HostKeyFile(s.config.HostKeyPath))
	}

	address := fmt.Sprintf(":%d", s.config.Port)
	return ssh.ListenAndServe(address, s.handleSession, serverOptions...)
}

func (s *Server) handleSession(session ssh.Session) {
	s.logger.WithField("user", session.User()).Info("New SSH session started")
	defer s.logger.WithField("user", session.User()).Info("SSH session ended")

	cleanScriptPath := filepath.Clean(s.config.ScriptPath)
	if _, err := os.Stat(cleanScriptPath); err != nil {
		s.logger.WithError(err).Error("Script file does not exist")
		return
	}

	cmd := exec.Command("/bin/bash", cleanScriptPath)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		s.logger.WithError(err).Error("Failed to start PTY")
		return
	}
	defer ptmx.Close()

	go func() {
		_, err := io.Copy(ptmx, session)
		if err != nil {
			s.logger.WithError(err).Error("Failed to copy input from SSH to PTY")
		}
	}()

	_, err = io.Copy(session, ptmx)
	if err != nil {
		s.logger.WithError(err).Error("Failed to copy output from PTY to SSH")
	}
}

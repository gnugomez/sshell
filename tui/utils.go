package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func CreateWindowSizeCmd(width, height int) tea.Cmd {
	return func() tea.Msg {
		return tea.WindowSizeMsg{Width: width, Height: height}
	}
}

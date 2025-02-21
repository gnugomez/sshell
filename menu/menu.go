package menu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
	choices  []string
	cursor   int
	selected int
	width    int
	height   int
}

func CreateMenu() *MenuModel {
	return &MenuModel{
		choices: []string{"About me", "Projects", "Blog"},
	}
}

func (m *MenuModel) Init() tea.Cmd { return nil }

func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m *MenuModel) View() string {

	titleStyle := lipgloss.NewStyle().Bold(true)
	footerStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("15")).
		Foreground(lipgloss.Color("0")).
		Padding(0, 1)

	var menu []string
	menu = append(menu, titleStyle.Render("Jordi Gómez\n"))

	for i, item := range m.choices {
		if m.cursor == i {
			menu = append(menu, "→ "+item)
		} else {
			menu = append(menu, "  "+item)
		}
	}

	// Main content (title + menu items)
	content := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height-1). // Reserve space for footer
		Align(lipgloss.Center, lipgloss.Center).
		Render(strings.Join(menu, "\n"))

	// Footer content
	footer := footerStyle.Render(
		"↑/↓: Navigate • Enter: Select • q: Quit",
	)

	// Combine both elements
	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Render(lipgloss.JoinVertical(
			content,
			footer,
		))
}

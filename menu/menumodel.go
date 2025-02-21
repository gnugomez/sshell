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
		choices: []string{"About me", "Blog", "Exit"},
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
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m *MenuModel) View() string {
	var menu []string

	titleStyle := lipgloss.NewStyle().Bold(true)
	menu = append(menu, titleStyle.Render("Jordi Gómez\n"))

	for i, item := range m.choices {
		if m.cursor == i {
			menu = append(menu, "→ "+item)
		} else {
			menu = append(menu, "  "+item)
		}
	}

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(strings.Join(menu, "\n"))
}

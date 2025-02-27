package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutModel struct {
	width  int
	height int
}

func CreateAbout() *AboutModel {
	return &AboutModel{}
}

func (m *AboutModel) Init() tea.Cmd { return nil }

func (m *AboutModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "esc":
			return CreateMenu(), CreateWindowSizeCmd(m.width, m.height)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m *AboutModel) View() string {
	// Define a maximum width for your content box.
	const maxContentWidth = 80

	titleStyle := lipgloss.NewStyle().Bold(true)
	textStyle := lipgloss.NewStyle().Align(lipgloss.Center)
	footerStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("15")).
		Foreground(lipgloss.Color("0")).
		Padding(0, 1)

	var text []string
	text = append(text, titleStyle.Render("About"))
	text = append(text, textStyle.Render("I'm a software engineer based in Barcelona, Spain. I'm passionate about open source software, distributed systems, and infrastructure automation."))

	// Determine the content box width: use the full width if the window is narrow,
	// or the fixed max width if the window is wider.
	contentWidth := m.width
	if m.width > maxContentWidth {
		contentWidth = maxContentWidth
	}

	// Create the content box with the determined width.
	contentBox := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center, lipgloss.Center).
		Padding(2).
		// Optionally, add a border if you want a visual box:
		// Border(lipgloss.NormalBorder()).
		Render(strings.Join(text, "\n"))

	// Center the content box within the available window space.
	centeredContent := lipgloss.Place(m.width, m.height-1, lipgloss.Center, lipgloss.Center, contentBox)

	// Render the footer.
	footer := footerStyle.Render("↑/↓: Navigate • Enter: Select • q: Quit • Esc: Back")

	// Combine both elements vertically.
	return lipgloss.JoinVertical(lipgloss.Top, centeredContent, footer)
}

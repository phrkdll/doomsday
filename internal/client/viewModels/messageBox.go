package viewModels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type MessageBox struct {
	Width   int
	Message string
}

func NewMessageBox(width int, message string) MessageBox {
	return MessageBox{
		Width:   width,
		Message: message,
	}
}

func (m MessageBox) Init() tea.Cmd {
	return nil
}

func (m MessageBox) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m MessageBox) View() string {
	w, h, _ := term.GetSize(0)
	style := lipgloss.NewStyle().
		Width(m.Width).
		Padding(1, 2).
		Margin((h-5)/2, (w-64)/2).
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFA500"))

	return style.Render(m.Message)
}

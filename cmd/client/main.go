package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/phrkdll/must/pkg/must"
)

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	must.Return(p.Run()).ElsePanic()
}

type model struct {
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return "Hello, Bubble Tea!"
}

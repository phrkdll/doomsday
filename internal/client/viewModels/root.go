package viewModels

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
)

type Root struct {
	Server string
}

func (m Root) Init() tea.Cmd {
	return nil
}

func (m Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":

			slog.Info("Gracefully exiting application", "key", msg.String())
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Root) View() string {
	message := "Welcome to Doomsday.\n\nA TUI-based multiplayer survival role-play/adventure game written in Go."
	return MessageBox{Width: 64, Message: message}.View()
}

package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/phrkdll/doomsday/internal/client/viewModels"
	"github.com/phrkdll/must/pkg/must"
)

func main() {
	f := must.Return(tea.LogToFile("debug.log", "debug")).ElsePanic()
	defer f.Close()

	p := tea.NewProgram(viewModels.Root{}, tea.WithAltScreen())
	must.Return(p.Run()).ElsePanic()
}

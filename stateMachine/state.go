package stateMachine

import tea "github.com/charmbracelet/bubbletea"

type State interface {
	Init()
	OnKeyPressed(key string) tea.Cmd
	Render() string
}

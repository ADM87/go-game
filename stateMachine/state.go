package stateMachine

import tea "github.com/charmbracelet/bubbletea"

type State interface {
	OnKeyPressed(key string) tea.Cmd
	Render() string
}

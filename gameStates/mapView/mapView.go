package mapView

import tea "github.com/charmbracelet/bubbletea"

type State struct {
}

func NewState() State {
	return State{}
}

func (s *State) OnKeyPressed(key string) tea.Cmd {
	return nil
}

func (s *State) Render() string {
	return "Map Display"
}

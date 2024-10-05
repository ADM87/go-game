package stateMachine

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	states  []State
	current int
}

func NewStateMachine(states []State, initial int) Model {
	return Model{states: states, current: initial}
}

func (m *Model) OnKeyPressed(key string) tea.Cmd {
	return m.states[m.current].OnKeyPressed(key)
}

func (m *Model) Render() string {
	return m.states[m.current].Render()
}

func (m *Model) SetState(index int) {
	m.current = index
}

func (m *Model) HasState(index int) bool {
	return index < len(m.states)
}

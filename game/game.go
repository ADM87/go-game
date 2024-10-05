package game

import (
	"go-game/gameStates/gameplay"
	"go-game/gameStates/mapView"
	"go-game/stateMachine"

	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	stateMachine stateMachine.Model
}

func NewGame() Model {
	gps := gameplay.NewState()
	mvs := mapView.NewState()
	return Model{
		stateMachine: stateMachine.NewStateMachine(
			[]stateMachine.State{
				&gps,
				&mvs,
			},
			0,
		),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "esc":
			return m, tea.Quit

		case "1", "2":
			s, e := strconv.Atoi(key)
			if e != nil {
				panic(e)
			}
			if m.stateMachine.HasState(s - 1) {
				m.stateMachine.SetState(s - 1)
			}

		default:
			return m, m.stateMachine.OnKeyPressed(key)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.stateMachine.Render()
}

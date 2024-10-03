package game

import (
	"go-game/geometry"

	game "go-game/game/states"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	StateMachine GameStateMachine
}

func NewGame() Model {
	overworldState := game.NewOverworldState(
		geometry.Point{X: 55, Y: 25},
		geometry.Point{X: 50, Y: 20},
	)

	states := map[string]game.GameState{
		"overworld": &overworldState,
	}

	return Model{
		StateMachine: NewGameStateMachine(states, "overworld"),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		default:
			return m, m.StateMachine.OnKeyPressed(msg.String())
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.StateMachine.Render()
}

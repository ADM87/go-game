package game

import (
	game "go-game/game/states"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	StateMachine GameStateMachine
}

func NewGame() Model {
	overworldState := game.NewOverworldState(50, 50, 35, 15)

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

		// Allow the game state machine to handle key presses
		default:
			return m, m.StateMachine.OnKeyPressed(msg.String())
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.StateMachine.Render()
}

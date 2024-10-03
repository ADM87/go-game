package game

import (
	game "go-game/game/states"

	tea "github.com/charmbracelet/bubbletea"
)

type GameStateMachine struct {
	States       map[string]game.GameState
	CurrentState string
}

func NewGameStateMachine(states map[string]game.GameState, initialState string) GameStateMachine {
	return GameStateMachine{States: states, CurrentState: initialState}
}

func (gsm *GameStateMachine) OnKeyPressed(key string) tea.Cmd {
	s := gsm.States[gsm.CurrentState]
	return s.OnKeyPressed(key)
}

func (gsm *GameStateMachine) Render() string {
	s := gsm.States[gsm.CurrentState]
	return s.Render()
}

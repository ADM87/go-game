package main

import (
	game "go-game/game"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(game.NewGame())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

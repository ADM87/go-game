package main

import (
	"go-game/game"
	"go-game/geometry"

	tea "github.com/charmbracelet/bubbletea"
)

var worldSize = geometry.Point{X: 55, Y: 25}
var viewSize = geometry.Point{X: 50, Y: 20}

func main() {
	p := tea.NewProgram(game.NewGame(worldSize, viewSize))
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

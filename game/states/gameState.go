package game

import tea "github.com/charmbracelet/bubbletea"

type GameState interface {
	OnKeyPressed(key string) tea.Cmd
	Render() string
}

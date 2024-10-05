package gameplay

import (
	tea "github.com/charmbracelet/bubbletea"
)

const _unknown int = -1
const _empty int = 0
const _wall int = 1
const _player int = 2

type State struct {
	world  World
	camera Camera
	tokens Tokens
	player Player
}

func NewState() State {
	w := NewWorld(100, 100)
	c := NewCamera(
		50,
		15)
	t := NewTokens(map[int]string{
		_unknown: "?",
		_empty:   " ",
		_wall:    "█",
		_player:  "☺",
	})
	p := NewPlayer(25, 25, _player, &w)
	c.Goto(p.Object.x, p.Object.y)
	return State{
		world:  w,
		camera: c,
		tokens: t,
		player: p,
	}
}

func (s *State) OnKeyPressed(key string) tea.Cmd {
	switch key {
	case "up":
		s.player.Move(0, -1)
		s.camera.Goto(s.player.Object.x, s.player.Object.y)

	case "down":
		s.player.Move(0, 1)
		s.camera.Goto(s.player.Object.x, s.player.Object.y)

	case "left":
		s.player.Move(-1, 0)
		s.camera.Goto(s.player.Object.x, s.player.Object.y)

	case "right":
		s.player.Move(1, 0)
		s.camera.Goto(s.player.Object.x, s.player.Object.y)
	}
	return nil
}

func (s *State) Render() string {
	return s.camera.Buffer(&s.world, &s.tokens)
}

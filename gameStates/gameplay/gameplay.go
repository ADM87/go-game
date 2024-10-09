package gameplay

import tea "github.com/charmbracelet/bubbletea"

const _unknown int = -1
const _empty int = 0
const _wall int = 1
const _player int = 2

type State struct {
	world  World
	camera Camera
	tokens Tokens
	player GameObject
}

func NewState() State {
	w := NewWorld(100, 100)
	c := NewCamera(50, 15)
	t := NewTokens(map[int]string{
		_unknown: " ",
		_empty:   " ",
		_wall:    "█",
		_player:  "☺",
	})
	p := NewGameObject(0, 0, _player, &w)
	return State{
		world:  w,
		camera: c,
		tokens: t,
		player: p,
	}
}

func (s *State) Init() tea.Cmd {
	s.player.SetPosition(5, 5)
	s.UpdateCamera()
	return nil
}

func (s *State) OnKeyPressed(key string) tea.Cmd {
	switch key {
	case "up", "down", "left", "right":
		switch key {
		case "up":
			s.player.Move(0, -1)
		case "down":
			s.player.Move(0, 1)
		case "left":
			s.player.Move(-1, 0)
		case "right":
			s.player.Move(1, 0)
		}
		s.UpdateCamera()
	}
	return nil
}

func (s *State) Render() string {
	return s.camera.Buffer(&s.world, &s.tokens)
}

func (s *State) UpdateCamera() {
	bx, by, bw, bh := s.world.Bounds()
	s.camera.Follow(&s.player.Entity)
	s.camera.BoundTo(bx, by, bx+bw, by+bh)
}

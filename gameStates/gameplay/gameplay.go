package gameplay

import (
	"go-game/data"

	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	world     World
	camera    Camera
	tokens    Tokens
	player    GameObject
	gameModel *data.GameModel
}

func NewState(mdl *data.GameModel) State {
	w := NewWorld(mdl.WorldWidth, mdl.WorldHeight)
	c := NewCamera(mdl.ViewWidth, mdl.ViewHeight)
	t := NewTokens(mdl.GameTokens)
	p := NewGameObject(0, 0, data.PlayerId, &w)
	return State{
		world:     w,
		camera:    c,
		tokens:    t,
		player:    p,
		gameModel: mdl,
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

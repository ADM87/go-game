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
	w := NewWorld()
	c := NewCamera(mdl.ViewWidth, mdl.ViewHeight)
	t := NewTokens(mdl.GameTokens)
	p := NewGameObject(-1, -1, data.PlayerId, &w)
	return State{
		world:     w,
		camera:    c,
		tokens:    t,
		player:    p,
		gameModel: mdl,
	}
}

func (s *State) Init() {
	s.player.SetPosition(
		s.world.CurrentMap().spawnX,
		s.world.CurrentMap().spawnY,
	)
	s.UpdateCamera()
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
	output := s.camera.Buffer(&s.world, &s.tokens)
	return output
}

func (s *State) UpdateCamera() {
	s.camera.Follow(&s.player.Entity)

	// If the world is larger than the camera view,  bound the camera to the edges
	if s.world.CurrentMap().width > s.camera.width {
		if s.camera.x < 0 {
			s.camera.x = 0
		}
		if s.camera.x > s.world.CurrentMap().width-s.camera.width {
			s.camera.x = s.world.CurrentMap().width - s.camera.width
		}
	}
	if s.world.CurrentMap().height > s.camera.height {
		if s.camera.y < 0 {
			s.camera.y = 0
		}
		if s.camera.y > s.world.CurrentMap().height-s.camera.height {
			s.camera.y = s.world.CurrentMap().height - s.camera.height
		}
	}
}

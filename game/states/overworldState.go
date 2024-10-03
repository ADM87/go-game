package game

import (
	"go-game/camera"
	"go-game/gameObject"
	"go-game/geometry"
	"go-game/world"

	tea "github.com/charmbracelet/bubbletea"
)

type OverworldState struct {
	Camera camera.Model
	World  world.Model
	Player gameObject.Model
}

func NewOverworldState(worldSize, viewSize geometry.Point) OverworldState {
	c := camera.NewCamera(viewSize.X, viewSize.Y) // Camera
	w := world.NewWorld(worldSize.X, worldSize.Y) // World
	p := gameObject.NewGameObject(1, 1, 1, &w)    // Player

	c.Follow(p.Position())
	c.ClampViewPort(geometry.Point{X: 0, Y: 0}, w.Size)

	return OverworldState{Camera: c, World: w, Player: p}
}

func (s *OverworldState) OnKeyPressed(key string) tea.Cmd {
	switch key {
	case "up", "w":
		movePlayer(0, -1, &s.Player, &s.Camera, &s.World)

	case "down", "s":
		movePlayer(0, 1, &s.Player, &s.Camera, &s.World)

	case "left", "a":
		movePlayer(-1, 0, &s.Player, &s.Camera, &s.World)

	case "right", "d":
		movePlayer(1, 0, &s.Player, &s.Camera, &s.World)
	}
	return nil
}

func (s *OverworldState) Render() string {
	return s.World.Render(s.Camera.ViewPort())
}

func movePlayer(x, y int, p *gameObject.Model, c *camera.Model, w *world.Model) {
	p.SafeMove(x, y)
	c.Follow(p.Position())
	c.ClampViewPort(geometry.Point{X: 0, Y: 0}, w.Size)
}

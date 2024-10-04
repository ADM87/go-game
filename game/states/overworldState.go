package game

import (
	"go-game/camera"
	"go-game/gameObject"
	"go-game/world"

	tea "github.com/charmbracelet/bubbletea"
)

type OverworldState struct {
	Camera camera.Model
	World  world.Model
	Player gameObject.Model
}

func NewOverworldState(worldWidth, worldHeight, viewWidth, viewHeight int) OverworldState {
	w := world.NewWorld(worldWidth, worldHeight) // World
	c := camera.NewCamera(viewWidth, viewHeight) // Camera
	p := gameObject.NewGameObject(0, 0, 1, &w)   // Player

	// Set the player's initial position
	movePlayer(5, 5, &p, &c, &w)

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
	if p.SafeMove(x, y) {
		c.Follow(p.Position())
		minX, minY := w.Min()
		maxX, maxY := w.Max()
		c.BoundTo(minX, minY, maxX, maxY)
	}
}

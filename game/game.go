package game

import (
	"go-game/camera"
	"go-game/geometry"
	"go-game/objects"
	"go-game/world"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Camera camera.Model
	World  world.Model
	Player objects.GameObject
}

func NewGame(worldSize, viewSize geometry.Point) Model {
	// Create critical game components
	c := camera.NewCamera(viewSize.X, viewSize.Y)
	w := world.NewWorld(worldSize.X, worldSize.Y)
	p := objects.NewGameObject(1, 1, 1, &w)

	// Set the camera's initial position to follow the player
	c.Follow(p.Position())

	// Return the game model
	return Model{Camera: c, World: w, Player: p}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		case "up", "w":
			movePlayer(0, -1, &m.Player, &m.Camera)

		case "down", "s":
			movePlayer(0, 1, &m.Player, &m.Camera)

		case "left", "a":
			movePlayer(-1, 0, &m.Player, &m.Camera)

		case "right", "d":
			movePlayer(1, 0, &m.Player, &m.Camera)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.World.Render(m.Camera.ViewPort())
}

func movePlayer(x, y int, p *objects.GameObject, c *camera.Model) {
	p.SafeMove(x, y)
	c.Follow(p.Position())
}

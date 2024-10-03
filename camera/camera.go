package camera

import (
	"go-game/geometry"
	"math"
)

type Model struct {
	viewPort geometry.Rectangle
}

func NewCamera(width, height int) Model {
	return Model{
		viewPort: geometry.Rectangle{
			Position: geometry.Point{
				X: -int(math.Floor(float64(width) / 2)),
				Y: -int(math.Floor(float64(height) / 2)),
			},
			Size: geometry.Point{X: width, Y: height},
		},
	}
}

func (c *Model) Follow(p geometry.Point) {
	c.viewPort.Position.X = p.X - int(math.Floor(float64(c.viewPort.Size.X)/2))
	c.viewPort.Position.Y = p.Y - int(math.Floor(float64(c.viewPort.Size.Y)/2))
}

func (c *Model) Move(x, y int) {
	c.viewPort.Position.X += x
	c.viewPort.Position.Y += y
}

func (c *Model) ViewPort() geometry.Rectangle {
	return c.viewPort
}

func (c *Model) ClampViewPort(min, max geometry.Point) {
	if c.viewPort.Position.X < min.X {
		c.viewPort.Position.X = min.X
	}
	if c.viewPort.Position.Y < min.Y {
		c.viewPort.Position.Y = min.Y
	}
	if c.viewPort.Position.X+c.viewPort.Size.X > max.X {
		c.viewPort.Position.X = max.X - c.viewPort.Size.X
	}
	if c.viewPort.Position.Y+c.viewPort.Size.Y > max.Y {
		c.viewPort.Position.Y = max.Y - c.viewPort.Size.Y
	}
}

package camera

import (
	"math"
)

type Model struct {
	x, y          int
	width, height int
	halfWidth     int
	halfHeight    int
}

func NewCamera(width, height int) Model {
	hw := int(math.Floor(float64(width) / 2))
	hh := int(math.Floor(float64(height) / 2))
	return Model{
		x:          -hw,
		y:          -hh,
		width:      width,
		height:     height,
		halfWidth:  hw,
		halfHeight: hh,
	}
}

func (c *Model) ViewPort() (int, int, int, int) {
	return c.x, c.y, c.x + c.width, c.y + c.height
}

func (c *Model) Follow(x, y int) {
	c.x = x - c.halfWidth
	c.y = y - c.halfHeight
}

func (c *Model) Move(x, y int) {
	c.x += x
	c.y += y
}

func (c *Model) BoundTo(minX, minY, maxX, maxY int) {
	if c.x < minX {
		c.x = minX
	}
	if c.y < minY {
		c.y = minY
	}
	if c.x > maxX-c.width {
		c.x = maxX - c.width
	}
	if c.y > maxY-c.height {
		c.y = maxY - c.height
	}
}

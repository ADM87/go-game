package gameplay

import "math"

type Camera struct {
	x, y          int
	width, height int
}

func NewCamera(width, height int) Camera {
	return Camera{
		x:      -int(math.Floor(float64(width) / 2)),
		y:      -int(math.Floor(float64(height) / 2)),
		width:  width,
		height: height,
	}
}

func (c *Camera) Follow(entity *Entity) {
	c.Goto(entity.x, entity.y)
}

func (c *Camera) Goto(x, y int) {
	c.x = x - int(math.Floor(float64(c.width)/2))
	c.y = y - int(math.Floor(float64(c.height)/2))
}

func (c *Camera) Move(x, y int) {
	c.x += x
	c.y += y
}

func (c *Camera) BoundTo(minX, minY, maxX, maxY int) {
	if c.x < minX {
		c.x = minX
	}
	if c.y < minY {
		c.y = minY
	}
	if c.x+c.width > maxX {
		c.x = maxX - c.width
	}
	if c.y+c.height > maxY {
		c.y = maxY - c.height
	}
}

func (c *Camera) Buffer(world *World, tokens *Tokens) string {
	output := ""
	for y := c.y; y < c.y+c.height; y++ {
		for x := c.x; x < c.x+c.width; x++ {
			output += tokens.GetToken(world.Get(x, y))
		}
		output += "\n"
	}
	return output
}

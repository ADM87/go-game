package gameObject

import (
	"go-game/world"
)

type Model struct {
	x, y  int
	tid   int
	world *world.Model
}

func NewGameObject(x, y, tid int, w *world.Model) Model {
	w.Set(x, y, tid)
	return Model{
		x:     x,
		y:     y,
		tid:   tid,
		world: w,
	}
}

func (g *Model) Position() (int, int) {
	return g.x, g.y
}

/* SetPosition sets the position of the GameObject to the given x and y coordinates without checking if the position is valid. */
func (g *Model) SetPosition(x, y int) {
	commitMove(x, y, g, g.world)
}

/* SafeSetPosition sets the position of the GameObject to the given x and y coordinates only if the position is valid. */
func (g *Model) SafeSetPosition(x, y int) bool {
	if canMoveTo(x, y, g.world) {
		commitMove(x, y, g, g.world)
		return true
	}
	return false
}

/* Move moves the GameObject by the given x and y coordinates without checking if the position is valid. */
func (g *Model) Move(x, y int) {
	commitMove(g.x+x, g.y+y, g, g.world)
}

/* SafeMove moves the GameObject by the given x and y coordinates only if the position is valid. */
func (g *Model) SafeMove(x, y int) bool {
	xx := g.x + x
	yy := g.y + y
	if canMoveTo(xx, yy, g.world) {
		commitMove(xx, yy, g, g.world)
		return true
	}
	return false
}

func canMoveTo(x, y int, w *world.Model) bool {
	minX, minY := w.Min()
	maxX, maxY := w.Max()
	if x < minX || y < minY || x >= maxX || y >= maxY {
		return false
	}
	return w.IsEmpty(x, y)
}

func commitMove(x, y int, g *Model, w *world.Model) {
	w.ClearAt(g.x, g.y)
	w.Set(x, y, g.tid)
	g.x = x
	g.y = y
}

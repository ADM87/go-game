package gameObject

import (
	"go-game/geometry"
	"go-game/world"
)

type Model struct {
	position geometry.Point
	tid      int
	world    *world.Model
}

func NewGameObject(x, y, tid int, w *world.Model) Model {
	w.Set(x, y, tid)
	return Model{
		position: geometry.Point{X: x, Y: y},
		tid:      tid,
		world:    w,
	}
}

func (g *Model) Position() geometry.Point {
	return g.position
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
	p := geometry.Point{X: g.position.X + x, Y: g.position.Y + y}
	commitMove(p.X, p.Y, g, g.world)
}

/* SafeMove moves the GameObject by the given x and y coordinates only if the position is valid. */
func (g *Model) SafeMove(x, y int) bool {
	p := geometry.Point{X: g.position.X + x, Y: g.position.Y + y}
	if canMoveTo(p.X, p.Y, g.world) {
		commitMove(p.X, p.Y, g, g.world)
		return true
	}
	return false
}

func canMoveTo(x, y int, w *world.Model) bool {
	if x < 0 || y < 0 || x >= w.Size.X || y >= w.Size.Y {
		return false
	}
	return w.IsEmpty(x, y)
}

func commitMove(x, y int, g *Model, w *world.Model) {
	w.ClearAt(g.position.X, g.position.Y)
	g.position.X = x
	g.position.Y = y
	w.Set(x, y, g.tid)
}

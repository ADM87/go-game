package objects

import (
	"go-game/geometry"
	"go-game/world"
)

type GameObject struct {
	position geometry.Point
	tid      int
	world    *world.Model
}

func NewGameObject(x, y, tid int, w *world.Model) GameObject {
	w.Set(x, y, tid)
	return GameObject{
		position: geometry.Point{X: x, Y: y},
		tid:      tid,
		world:    w,
	}
}

func (g *GameObject) Position() geometry.Point {
	return g.position
}

/* SetPosition sets the position of the GameObject to the given x and y coordinates without checking if the position is valid. */
func (g *GameObject) SetPosition(x, y int) {
	commitMove(x, y, g, g.world)
}

/* SafeSetPosition sets the position of the GameObject to the given x and y coordinates only if the position is valid. */
func (g *GameObject) SafeSetPosition(x, y int) {
	if canMoveTo(x, y, g.world) {
		commitMove(x, y, g, g.world)
	}
}

/* Move moves the GameObject by the given x and y coordinates without checking if the position is valid. */
func (g *GameObject) Move(x, y int) {
	p := geometry.Point{X: g.position.X + x, Y: g.position.Y + y}
	commitMove(p.X, p.Y, g, g.world)
}

/* SafeMove moves the GameObject by the given x and y coordinates only if the position is valid. */
func (g *GameObject) SafeMove(x, y int) {
	p := geometry.Point{X: g.position.X + x, Y: g.position.Y + y}
	if canMoveTo(p.X, p.Y, g.world) {
		commitMove(p.X, p.Y, g, g.world)
	}
}

func canMoveTo(x, y int, w *world.Model) bool {
	if x < 0 || y < 0 || x >= w.Size.X || y >= w.Size.Y {
		return false
	}
	return w.IsEmpty(x, y)
}

func commitMove(x, y int, g *GameObject, w *world.Model) {
	w.Set(g.position.X, g.position.Y, 0)
	g.position.X = x
	g.position.Y = y
	w.Set(x, y, g.tid)
}

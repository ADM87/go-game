package gameplay

/* Represents a position and token within the game world. */
type Entity struct {
	x, y int
	id   int
}

func NewEntity(x, y, id int) Entity {
	return Entity{x: x, y: y, id: id}
}

/* Represents an entity that can move within the game world. */
type GameObject struct {
	Entity Entity
	World  *World
}

func NewGameObject(x, y, id int, world *World) GameObject {
	world.Set(x, y, id)
	return GameObject{
		Entity: NewEntity(x, y, id),
		World:  world,
	}
}

/* Moves the object by dx and dy if the new position is empty. Returns whethers or not the move was successful. */
func (o *GameObject) Move(dx, dy int) bool {
	if o.World.IsEmpty(o.Entity.x+dx, o.Entity.y+dy) {
		o.World.Set(o.Entity.x, o.Entity.y, _empty)
		o.Entity.x += dx
		o.Entity.y += dy
		o.World.Set(o.Entity.x, o.Entity.y, o.Entity.id)
		return true
	}
	return false
}

/* Sets the object's position to x and y if the new position is empty. Returns whether or not the set was successful. */
func (o *GameObject) SetPosition(x, y int) bool {
	if o.World.IsEmpty(x, y) {
		o.World.Set(o.Entity.x, o.Entity.y, _empty)
		o.Entity.x = x
		o.Entity.y = y
		o.World.Set(o.Entity.x, o.Entity.y, o.Entity.id)
		return true
	}
	return false
}

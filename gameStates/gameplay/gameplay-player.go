package gameplay

type Player struct {
	Object Object
	World  *World
}

func NewPlayer(x, y, id int, world *World) Player {
	world.Set(x, y, id)
	return Player{
		Object: NewObject(x, y, id),
		World:  world,
	}
}

func (p *Player) Move(dx, dy int) {
	if p.World.IsEmpty(p.Object.x+dx, p.Object.y+dy) {
		p.World.Set(p.Object.x, p.Object.y, _empty)
		p.Object.x += dx
		p.Object.y += dy
		p.World.Set(p.Object.x, p.Object.y, p.Object.id)
	}
}

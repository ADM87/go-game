package gameplay

import (
	"go-game/data"
	"math"
	"math/rand"
)

const (
	MinRoomCount  = 50
	MaxRoomCount  = 60
	MinRoomWidth  = 11
	MaxRoomWidth  = 13
	MinRoomHeight = 7
	MaxRoomHeight = 9
)

const (
	North = iota
	East
	South
	West
)

type Room struct {
	width, height int
	x, y          int
	neighbors     []*Room
}

func NewRoom(w, h int) Room {
	return Room{
		width:     w,
		height:    h,
		x:         0,
		y:         0,
		neighbors: make([]*Room, 4),
	}
}

func (r *Room) MinX() int {
	return r.x
}

func (r *Room) MinY() int {
	return r.y
}

func (r *Room) Min() (int, int) {
	return r.MinX(), r.MinY()
}

func (r *Room) MaxX() int {
	return r.x + r.width
}

func (r *Room) MaxY() int {
	return r.y + r.height
}

func (r *Room) Max() (int, int) {
	return r.MaxX(), r.MaxY()
}

func (r *Room) CenterX() int {
	return r.x + int(math.Floor(float64(r.width/2)))
}

func (r *Room) CenterY() int {
	return r.y + int(math.Floor(float64(r.height/2)))
}

func (r *Room) Center() (int, int) {
	return r.CenterX(), r.CenterY()
}

type Map struct {
	spawnX, spawnY int
	exitX, exitY   int
	width, height  int
	rooms          []Room
	layout         []int
}

func NewMap() Map {
	newMap := Map{}

	w, h := getRoomSize()

	// Create the first room
	r := NewRoom(w, h)

	// Add the first room to the map
	newMap.rooms = append(newMap.rooms, r)

	maxRooms := max(MinRoomCount, MaxRoomCount)
	if MaxRoomCount != MinRoomCount {
		maxRooms = MinRoomCount + rand.Intn(MaxRoomCount-MinRoomCount)
	}

	// Generate rooms
	generateRooms(&r, &newMap.rooms, maxRooms)

	// Fix room coordinates so they are not negative
	fixRoomCoordinates(&newMap.rooms)

	// Calculate the map's dimensions
	newMap.width, newMap.height = calculateMapDimensions(&newMap.rooms)

	// Define layout
	newMap.layout = make([]int, newMap.width*newMap.height)

	// Draw layout
	drawWalls(&newMap.rooms, &newMap.layout, newMap.width)
	drawDoors(&newMap.rooms, &newMap.layout, newMap.width)

	// Set spawn and exit
	newMap.spawnX, newMap.spawnY = newMap.rooms[0].Center()
	newMap.exitX, newMap.exitY = newMap.rooms[len(newMap.rooms)-1].Center()

	return newMap
}

func getRoomSize() (int, int) {
	w := MinRoomWidth + rand.Intn(MaxRoomWidth-MinRoomWidth)
	h := MinRoomHeight + rand.Intn(MaxRoomHeight-MinRoomHeight)
	if w%2 == 0 {
		w++
	}
	if h%2 == 0 {
		h++
	}
	return w, h
}

func generateRooms(cr *Room, rs *[]Room, max int) {
	if len(*rs) == max {
		return
	}

	directions := []int{North, East, South, West}

	for {
		if len(directions) == 0 || len(*rs) == max {
			return
		}

		i := rand.Intn(len(directions))
		dir := directions[i]
		directions = append(directions[:i], directions[i+1:]...)

		if cr.neighbors[dir] != nil {
			generateRooms(cr.neighbors[dir], rs, max)
			continue
		}

		w, h := getRoomSize()
		nr := NewRoom(w, h)

		switch dir {
		case North:
			nr.x = cr.CenterX() - int(math.Floor(float64(w/2)))
			nr.y = cr.MinY() - h
		case South:
			nr.x = cr.CenterX() - int(math.Floor(float64(w/2)))
			nr.y = cr.MaxY()
		case East:
			nr.x = cr.MaxX()
			nr.y = cr.CenterY() - int(math.Floor(float64(h/2)))
		case West:
			nr.x = cr.MinX() - w
			nr.y = cr.CenterY() - int(math.Floor(float64(h/2)))
		}

		isOverlapping := false
		for i := range *rs {
			if nr.x < (*rs)[i].MaxX() && nr.MaxX() > (*rs)[i].MinX() && nr.y < (*rs)[i].MaxY() && nr.MaxY() > (*rs)[i].MinY() {
				isOverlapping = true
				break
			}
		}
		if isOverlapping {
			continue
		}

		cr.neighbors[dir] = &(*rs)[len(*rs)-1]
		nr.neighbors[(dir+2)%4] = cr

		*rs = append(*rs, nr)

		generateRooms(&nr, rs, max)
	}
}

func fixRoomCoordinates(r *[]Room) {
	minX, minY := 0, 0
	for i := range *r {
		minX = min(minX, (*r)[i].MinX())
		minY = min(minY, (*r)[i].MinY())
	}
	for i := range *r {
		(*r)[i].x -= minX
		(*r)[i].y -= minY
	}
}

func calculateMapDimensions(r *[]Room) (int, int) {
	minX, minY := 0, 0
	maxX, maxY := 0, 0
	for i := range *r {
		minX = min(minX, (*r)[i].MinX())
		minY = min(minY, (*r)[i].MinY())
	}
	for i := range *r {
		maxX = max(maxX, (*r)[i].MaxX())
		maxY = max(maxY, (*r)[i].MaxY())
	}
	return maxX - minX, maxY - minY
}

func drawWalls(r *[]Room, l *[]int, w int) {
	for i := range *r {
		for x := (*r)[i].MinX(); x < (*r)[i].MaxX(); x++ {
			for y := (*r)[i].MinY(); y < (*r)[i].MaxY(); y++ {
				if x == (*r)[i].MinX() || x == (*r)[i].MaxX()-1 || y == (*r)[i].MinY() || y == (*r)[i].MaxY()-1 {
					(*l)[x+y*w] = data.WallId
				}
			}
		}
	}
}

func drawDoors(r *[]Room, l *[]int, w int) {
	for i := range *r {
		for j := range (*r)[i].neighbors {
			if (*r)[i].neighbors[j] == nil {
				continue
			}

			x, y := (*r)[i].Center()
			minX, minY := (*r)[i].Min()
			maxX, maxY := (*r)[i].Max()
			minX -= 1
			minY -= 1

			switch j {
			case North:
				for k := x - 2; k <= x+2; k++ {
					(*l)[k+minY*w] = data.EmptyId
				}
			case South:
				for k := x - 2; k <= x+2; k++ {
					(*l)[k+maxY*w] = data.EmptyId
				}
			case East:
				for k := y - 1; k <= y+1; k++ {
					(*l)[maxX+k*w] = data.EmptyId
				}
			case West:
				for k := y - 1; k <= y+1; k++ {
					(*l)[minX+k*w] = data.EmptyId
				}
			}
		}
	}
}

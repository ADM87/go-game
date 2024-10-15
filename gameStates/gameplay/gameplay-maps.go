package gameplay

import (
	"golang.org/x/exp/rand"
)

const (
	MinRoomCount        = 15
	MaxRoomCount        = 20
	MinRoomWidth        = 15
	MaxRoomWidth        = 20
	MinRoomHeight       = 5
	MaxRoomHeight       = 10
	PaddingBetweenRooms = 5
)

type Room struct {
	minX, minY int
	maxX, maxY int
	x, y       int
	neighbors  []*Room
}

type Map struct {
	spawn  int     // Index of the spawn point
	exit   int     // Index of the exit point
	rooms  []*Room // List of rooms
	layout []int   // 1D array representing the map layout
}

func NewMap(width, height int) Map {
	rs := make([]*Room, 0)
	rx := width / 2
	ry := height / 2
	rw := MinRoomWidth + rand.Intn(MaxRoomWidth-MinRoomWidth)
	rh := MinRoomHeight + rand.Intn(MaxRoomHeight-MinRoomHeight)
	r := Room{
		minX:      rx - rw/2,
		minY:      ry - rh/2,
		maxX:      rx + rw/2,
		maxY:      ry + rh/2,
		x:         rx,
		y:         ry,
		neighbors: make([]*Room, 4),
	}
	rs = append(rs, &r)
	roomCount := MinRoomCount + rand.Intn(MaxRoomCount-MinRoomCount)

	generate(&r, &rs, roomCount, width, height)
	return Map{
		spawn:  rs[0].y*width + rs[0].x,
		exit:   0,
		rooms:  rs,
		layout: makeLayout(&rs, width, height),
	}
}

func generate(r *Room, rs *[]*Room, rc, w, h int) (bool, *[]*Room) {
	if len(*rs) >= rc {
		return false, rs
	}

	rw := MinRoomWidth + rand.Intn(MaxRoomWidth-MinRoomWidth)
	rh := MinRoomHeight + rand.Intn(MaxRoomHeight-MinRoomHeight)
	hw := rw / 2
	hh := rh / 2

	// North, East, South, West
	directions := []int{0, 1, 2, 3}

	for {
		if len(directions) == 0 || len(*rs) >= rc {
			return false, rs
		}

		i := rand.Intn(len(directions))
		dir := directions[i]
		// Remove the direction from the list
		directions = append(directions[:i], directions[i+1:]...)

		if r.neighbors[dir] != nil {
			continue
		}

		var nx, ny int
		switch dir {
		case 0: // North
			nx = r.x
			ny = r.y - rh - PaddingBetweenRooms
		case 1: // East
			nx = r.x + rw + PaddingBetweenRooms
			ny = r.y
		case 2: // South
			nx = r.x
			ny = r.y + rh + PaddingBetweenRooms
		case 3: // West
			nx = r.x - rw - PaddingBetweenRooms
			ny = r.y
		}

		// Create the new room
		nr := Room{
			minX:      nx - hw,
			minY:      ny - hh,
			maxX:      nx + hw,
			maxY:      ny + hh,
			x:         nx,
			y:         ny,
			neighbors: make([]*Room, 4),
		}

		// Check if the room is within bounds
		if nr.minX < 0 || nr.minY < 0 || nr.maxX >= w || nr.maxY >= h {
			continue
		}

		// Check if the room overlaps with any existing rooms
		overlaps := false
		for i := 0; i < len(*rs); i++ {
			if nr.minX <= (*rs)[i].maxX && nr.maxX >= (*rs)[i].minX &&
				nr.minY <= (*rs)[i].maxY && nr.maxY >= (*rs)[i].minY {
				overlaps = true
				break
			}
		}
		if overlaps {
			continue
		}

		// Set the neighbors: N<>S, E<>W
		// Set the neighbors: N<>S, E<>W
		oppositeDir := (dir + 2) % 4
		r.neighbors[dir] = &nr
		nr.neighbors[oppositeDir] = r

		// Add the room to the list
		*rs = append(*rs, &nr)

		generate(&nr, rs, rc, w, h)
	}
}

func makeLayout(rs *[]*Room, w, h int) []int {
	l := make([]int, w*h)

	// Draw room outlines
	for i := 0; i < len(*rs); i++ {
		for y := (*rs)[i].minY; y <= (*rs)[i].maxY; y++ {
			for x := (*rs)[i].minX; x <= (*rs)[i].maxX; x++ {
				if x == (*rs)[i].minX || x == (*rs)[i].maxX || y == (*rs)[i].minY || y == (*rs)[i].maxY {
					l[y*w+x] = 1
				}
			}
		}

		// Connect neighbors
		for j := 0; j < len((*rs)[i].neighbors); j++ {
			room := (*rs)[i].neighbors[j]

			if room == nil {
				continue
			}

			// Clear doorways
			switch j {
			case 0: // North
				l[(*rs)[i].minY*w+(*rs)[i].x-1] = 0
				l[(*rs)[i].minY*w+(*rs)[i].x] = 0
				l[(*rs)[i].minY*w+(*rs)[i].x+1] = 0
				for y := (*rs)[i].minY - 1; y >= (*rs)[i].neighbors[j].maxY; y-- {
					l[y*w+(*rs)[i].x-2] = 1
					l[y*w+(*rs)[i].x+2] = 1
				}
			case 1: // East
				l[((*rs)[i].y-1)*w+(*rs)[i].maxX] = 0
				l[(*rs)[i].y*w+(*rs)[i].maxX] = 0
				l[((*rs)[i].y+1)*w+(*rs)[i].maxX] = 0
				for x := (*rs)[i].maxX + 1; x <= (*rs)[i].neighbors[j].minX; x++ {
					l[((*rs)[i].y-2)*w+x] = 1
					l[((*rs)[i].y+2)*w+x] = 1
				}
			case 2: // South
				l[(*rs)[i].maxY*w+(*rs)[i].x-1] = 0
				l[(*rs)[i].maxY*w+(*rs)[i].x] = 0
				l[(*rs)[i].maxY*w+(*rs)[i].x+1] = 0
			case 3: // West
				l[((*rs)[i].y-1)*w+(*rs)[i].minX] = 0
				l[(*rs)[i].y*w+(*rs)[i].minX] = 0
				l[((*rs)[i].y+1)*w+(*rs)[i].minX] = 0
			}
		}
	}

	return l
}

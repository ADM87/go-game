package gameplay

import "go-game/data"

type World struct {
	width, height int
	matrix        []int
}

func NewWorld(width, height int) World {
	return World{
		width:  width,
		height: height,
		matrix: make([]int, width*height),
	}
}

func (w *World) Set(x, y, id int) {
	if x < 0 || x >= w.width || y < 0 || y >= w.height {
		return
	}
	w.matrix[y*w.width+x] = id
}

func (w *World) Get(x, y int) int {
	if x < 0 || x >= w.width || y < 0 || y >= w.height {
		return data.UnknownId
	}
	return w.matrix[y*w.width+x]
}

func (w *World) Bounds() (int, int, int, int) {
	return 0, 0, w.width, w.height
}

func (w *World) IsEmpty(x, y int) bool {
	return w.Get(x, y) == data.EmptyId
}

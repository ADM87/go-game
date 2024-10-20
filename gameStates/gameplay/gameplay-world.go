package gameplay

import "go-game/data"

type World struct {
	maps       []Map
	currentMap int
}

func NewWorld() World {
	m := make([]Map, 1)
	m[0] = NewMap()
	return World{
		maps: m,
	}
}

func (w *World) CurrentMap() *Map {
	return &w.maps[w.currentMap]
}

func (w *World) Width() int {
	return w.CurrentMap().width
}

func (w *World) Height() int {
	return w.CurrentMap().height
}

func (w *World) Set(x, y, id int) {
	if w.maps[w.currentMap].width == 0 || w.maps[w.currentMap].height == 0 || x < 0 ||
		x >= w.maps[w.currentMap].width || y < 0 || y >= w.maps[w.currentMap].height {
		return
	}
	w.maps[w.currentMap].layout[y*w.maps[w.currentMap].width+x] = id
}

func (w *World) Get(x, y int) int {
	if x < 0 || x >= w.maps[w.currentMap].width || y < 0 || y >= w.maps[w.currentMap].height {
		return data.UnknownId
	}
	return w.maps[w.currentMap].layout[y*w.maps[w.currentMap].width+x]
}

func (w *World) Bounds() (int, int, int, int) {
	return 0, 0, w.maps[w.currentMap].width, w.maps[w.currentMap].height
}

func (w *World) IsEmpty(x, y int) bool {
	return w.Get(x, y) == data.EmptyId
}

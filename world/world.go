package world

import "go-game/geometry"

type Model struct {
	matrix []int
	Size   geometry.Point
}

func NewWorld(width, height int) Model {
	return Model{
		matrix: make([]int, width*height),
		Size:   geometry.Point{X: width, Y: height},
	}
}

func (w Model) Get(x, y int) int {
	return w.matrix[y*w.Size.X+x]
}

func (w *Model) Set(x, y, value int) {
	w.matrix[y*w.Size.X+x] = value
}

func (w *Model) IsEmpty(x, y int) bool {
	return w.Get(x, y) == 0
}

// Render returns a string representation of the world within the given viewPort.
func (w *Model) Render(viewPort geometry.Rectangle) string {
	output := ""
	for y := viewPort.Position.Y; y < viewPort.Position.Y+viewPort.Size.Y; y++ {
		for x := viewPort.Position.X; x < viewPort.Position.X+viewPort.Size.X; x++ {
			if x < 0 || y < 0 || x >= w.Size.X || y >= w.Size.Y {
				output += "?" // Out of bounds
			} else {
				switch w.Get(x, y) {
				case 0:
					output += " "
				case 1:
					output += "X"
				}
			}
		}
		output += "\n"
	}
	return output
}

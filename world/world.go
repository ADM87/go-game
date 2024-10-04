package world

type Model struct {
	matrix        []int
	width, height int
}

func NewWorld(width, height int) Model {
	return Model{
		matrix: make([]int, width*height),
		width:  width,
		height: height,
	}
}

func (w *Model) Min() (x, y int) {
	return 0, 0
}

func (w *Model) Max() (x, y int) {
	return w.width, w.height
}

func (w *Model) Get(x, y int) int {
	return w.matrix[y*w.width+x]
}

func (w *Model) Set(x, y, value int) {
	w.matrix[y*w.width+x] = value
}

func (w *Model) ClearAt(x, y int) {
	w.Set(x, y, 0)
}

func (w *Model) IsEmpty(x, y int) bool {
	return w.Get(x, y) == 0
}

// Render returns a string representation of the world within the given viewPort.
func (w *Model) Render(minX, minY, maxX, maxY int) string {
	output := ""
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			if x < 0 || y < 0 || x >= w.width || y >= w.height {
				output += "?" // Out of boundsw.height
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

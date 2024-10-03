package geometry

/*****************
 * Point
 *****************/
type Point struct {
	X, Y int
}

func (p Point) AddPoint(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) AddXY(x, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

/*****************
 * Rectangle
 *****************/
type Rectangle struct {
	Position Point
	Size     Point
}

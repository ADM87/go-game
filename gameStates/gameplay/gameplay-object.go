package gameplay

type Object struct {
	x, y int
	id   int
}

func NewObject(x, y, id int) Object {
	return Object{x: x, y: y, id: id}
}

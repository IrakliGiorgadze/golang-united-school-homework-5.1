package square

type Point struct {
	x int
	y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	return Point{
		s.start.x + int(s.a),
		s.start.y + int(s.a),
	}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}

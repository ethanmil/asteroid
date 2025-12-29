package physics

type Direction string

const (
	First  Direction = "first"
	Second Direction = "second"
	Third  Direction = "third"
	Fourth Direction = "fourth"
)

type Size struct {
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
}

type Position Point

func (p *Position) GetCoordinates(scale int) (x, y int) {
	return p.X / scale, p.Y / scale
}

func (p *Position) Offset(pos *Position) *Position {
	return &Position{X: p.X - pos.X, Y: p.Y - pos.Y}
}

func (p *Position) OffsetWithSize(s *Size) *Position {
	return &Position{X: p.X + s.Width, Y: p.Y + s.Height}
}

type Vector Point

func GetDirection(currentPosition, nextPosition Position) Direction {
	var (
		deltaX = nextPosition.X - currentPosition.X
		deltaY = nextPosition.Y - currentPosition.Y
	)

	if deltaX <= 0 && deltaY <= 0 {
		return Second
	} else if deltaX >= 0 && deltaY <= 0 {
		return First
	} else if deltaX <= 0 && deltaY >= 0 {
		return Third
	} else {
		return Fourth
	}
}

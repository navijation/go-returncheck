package point

type Point struct {
	X int
	Y int
}

func (me Point) SetX(x int) Point {
	me.X = x
	return me
}

func (me Point) SetY(y int) Point {
	me.Y = y
	return me
}

func (me Point) Copy() Point {
	return me
}

func (me Point) GetX() int {
	return me.X
}

func (me Point) GetY() int {
	return me.Y
}

func (me Point) Unpack() (int, int) {
	return me.X, me.Y
}

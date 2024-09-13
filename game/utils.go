package game

const (
	screenWidth  = 800
	screenHeight = 600
)

type Vector struct {
	X float64
	Y float64
}

type DivRectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewDivRectangle(x, y, width, height float64) DivRectangle {

	return DivRectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}

}

func (d DivRectangle) Intersects(other DivRectangle) bool {
	return d.X <= other.maxX() && other.X <= d.maxX() && d.Y <= other.maxY() && other.Y <= d.maxY()
}

func (d DivRectangle) maxX() float64 {
	return d.X + d.Width
}

func (d DivRectangle) maxY() float64 {
	return d.Y + d.Height
}

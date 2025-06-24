package areacalc

import "strings"

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	a     float64
	b     float64
	shape string
}

func NewRectangle(a float64, b float64, shape string) *Rectangle {
	return &Rectangle{
		a:     a,
		b:     b,
		shape: shape,
	}
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.a * rectangle.b
}

func (rectangle Rectangle) Type() string {
	return rectangle.shape
}

type Circle struct {
	r     float64
	shape string
}

func NewCircle(r float64, shape string) *Circle {
	return &Circle{
		r:     r,
		shape: shape,
	}
}

func (circle Circle) Area() float64 {
	return circle.r * circle.r * pi
}

func (circle Circle) Type() string {
	return circle.shape
}

func AreaCalculator(figures []Shape) (string, float64) {
	var (
		nameOfFigures []string
		totalArea     float64
	)

	for _, fig := range figures {
		nameOfFigures = append(nameOfFigures, fig.Type())
		totalArea += fig.Area()
	}

	return strings.Join(nameOfFigures, "-"), totalArea
}

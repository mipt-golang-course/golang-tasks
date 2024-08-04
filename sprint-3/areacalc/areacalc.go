package areacalc

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	width, height float64
	name          string
}

func NewRectangle(width, height float64, name string) *Rectangle {
	return &Rectangle{width: width, height: height, name: name}
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (Rectangle) Type() string {
	return "rectangle"
}

type Circle struct {
	radius float64
	name   string
}

func NewCircle(radius float64, name string) *Circle {
	return &Circle{radius: radius, name: name}
}

func (cir Circle) Area() float64 {
	return pi * cir.radius * cir.radius
}

func (Circle) Type() string {
	return "circle"
}

func AreaCalculator(figures []Shape) (string, float64) {
	if len(figures) == 0 {
		return "", 0.0
	}

	ret := 0.0

	var str string

	ret += figures[0].Area()
	str += figures[0].Type()

	for _, fig := range figures[1:] {
		ret += fig.Area()
		str += "-"
		str += fig.Type()
	}

	return str, ret

}

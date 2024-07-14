package areacalc

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	// TODO: implement me
}

func NewRectangle(float64, float64, string) *Rectangle {
	// TODO: implement me
	return &Rectangle{}
}

// TODO: implement me

type Circle struct {
	// TODO: implement me
}

func NewCircle(float64, string) *Circle {
	// TODO: implement me
	return &Circle{}
}

func AreaCalculator(figures []Shape) (string, float64) {
	// TODO: implement me
	return "", 0.0
}

package areacalc

import "testing"

const (
	shapeRectangle = "rectangle"
	shapeCircle    = "circle"
)

func TestAreaCalculator(t *testing.T) {
	tests := []struct {
		name    string
		figures []Shape
		want    string
		want1   float64
	}{
		{
			"test rectangle-rectangle",
			[]Shape{NewRectangle(1, 1, shapeRectangle), NewRectangle(2, 2, shapeRectangle)},
			"rectangle-rectangle",
			5,
		},
		{
			"test rectangle-circle-rectangle",
			[]Shape{NewRectangle(1, 1, shapeRectangle), NewCircle(1, shapeCircle), NewRectangle(2, 2, shapeRectangle)},
			"rectangle-circle-rectangle",
			8.14159,
		},
		{
			"empty",
			nil,
			"",
			0,
		},
		{
			"test olympics",
			[]Shape{NewCircle(1, shapeCircle), NewCircle(1, shapeCircle), NewCircle(1, shapeCircle), NewCircle(1, shapeCircle), NewCircle(1, shapeCircle)},
			"circle-circle-circle-circle-circle",
			15.70795,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AreaCalculator(tt.figures)
			if got != tt.want {
				t.Errorf("AreaCalculator() got = %v, want %v", got, tt.want)
			}

			if got1 != tt.want1 {
				t.Errorf("AreaCalculator() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

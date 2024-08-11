package areacalc_test

import (
	"testing"

	calc "github.com/mipt-golang-course/golang-tasks/sprint-3/areacalc"
)

const (
	shapeRectangle = "rectangle"
	shapeCircle    = "circle"
)

func TestAreaCalculator(t *testing.T) {
	tests := []struct {
		name    string
		figures []calc.Shape
		want    string
		want1   float64
	}{
		{
			"test rectangle-rectangle",
			[]calc.Shape{calc.NewRectangle(1, 1, shapeRectangle), calc.NewRectangle(2, 2, shapeRectangle)},
			"rectangle-rectangle",
			5,
		},
		{
			"test rectangle-circle-rectangle",
			[]calc.Shape{calc.NewRectangle(1, 1, shapeRectangle), calc.NewCircle(1, shapeCircle), calc.NewRectangle(2, 2, shapeRectangle)},
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
			[]calc.Shape{calc.NewCircle(1, shapeCircle), calc.NewCircle(1, shapeCircle), calc.NewCircle(1, shapeCircle), calc.NewCircle(1, shapeCircle), calc.NewCircle(1, shapeCircle)},
			"circle-circle-circle-circle-circle",
			15.70795,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calc.AreaCalculator(tt.figures)
			if got != tt.want {
				t.Errorf("calc.AreaCalculator() got = %v, want %v", got, tt.want)
			}

			if got1 != tt.want1 {
				t.Errorf("calc.AreaCalculator() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

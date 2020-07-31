package go_learning

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle {Width:12, Height:6}, want:72.0},
		{shape: Circle {Radius: 10}, want: 314.1592653589793},
		{shape: Triangle {Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64  {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width float64
	Height float64
}
func (r Rectangle) Area() float64  {
	return r.Width * r.Height
}

type Triangle struct {
	Base float64
	Height float64
}
func(t Triangle)  Area() float64 {
	return (t.Base * t.Height) / 2
}

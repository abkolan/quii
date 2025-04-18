package main

import "testing"

func TestAreaTableDrivenTests(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g but want %g", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if want != got {
			t.Errorf("got %g but want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

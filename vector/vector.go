package vector

import "math"

type Vector2D struct {
	X, Y float64
}

func (v *Vector2D) Add(vector *Vector2D) *Vector2D {
	return &Vector2D{X: v.X + vector.X, Y: v.Y + vector.Y}
}

func (v *Vector2D) Subtract(vector *Vector2D) *Vector2D {
	return &Vector2D{X: v.X - vector.X, Y: v.Y - vector.Y}
}

func (v *Vector2D) Multiply(vector *Vector2D) *Vector2D {
	return &Vector2D{X: v.X * vector.X, Y: v.Y * vector.Y}
}

func (v *Vector2D) AddV(d float64) *Vector2D {
	return &Vector2D{X: v.X + d, Y: v.Y + d}
}

func (v *Vector2D) MultiplyV(d float64) *Vector2D {
	return &Vector2D{X: v.X * d, Y: v.Y * d}
}

func (v *Vector2D) DivisionV(d float64) *Vector2D {
	return &Vector2D{X: v.X / d, Y: v.Y / d}
}

func (v *Vector2D) Limit(lower, upper float64) *Vector2D {
	return &Vector2D{
		X: math.Min(math.Max(v.X, lower), upper),
		Y: math.Min(math.Max(v.Y, lower), upper),
	}
}

func (v *Vector2D) Distance(vector *Vector2D) float64 {
	return math.Sqrt(math.Pow(v.X-vector.X, 2) + math.Pow(v.Y-vector.Y, 2))
}

package vector

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func New(x, y float64) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

func (vec Vector2) Mag() float64 {
	return math.Sqrt((float64(vec.X) * float64(vec.X)) + (float64(vec.Y) * float64(vec.Y)))
}

func (vec1 *Vector2) Add(vec2 Vector2) {
	vec1.X += vec2.X
	vec1.Y += vec2.Y
}

func (vec1 *Vector2) Sub(vec2 Vector2) {
	vec1.X -= vec2.X
	vec1.Y -= vec2.Y
}

func (vec Vector2) Mult(n float64) Vector2 {
	vec.X *= n
	vec.Y *= n

	return vec
}

func (vec *Vector2) Dot(vec2 Vector2) float64 {
	return vec2.X*vec2.X + vec.Y*vec2.Y
}

func (vec *Vector2) Normal() Vector2 {
	return Vector2{
		X: vec.Y,
		Y: -vec.X,
	}
}

package vector

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

func (vec1 Vector2) Add(vec2 Vector2) Vector2 {
	return Vector2{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
}

func (vec1 Vector2) Sub(vec2 Vector2) Vector2 {
	vec1.X -= vec2.X
	vec1.Y -= vec2.Y

	return vec1
}

func (vec Vector2) Scale(n float64) Vector2 {
	return Vector2{
		X: vec.X * n,
		Y: vec.Y * n,
	}
}

func (vec Vector2) Rotate(angle float64) Vector2 {
	result := Vector2{}
	result.X = vec.X*math.Cos(angle) - vec.Y*math.Sin(angle)
	result.Y = vec.X*math.Sin(angle) + vec.Y*math.Cos(angle)

	return result
}

func (vec Vector2) Mag() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

func (vec Vector2) MagSqr() float64 {
	return vec.X*vec.X + vec.Y*vec.Y
}

func (vec Vector2) Normalize() Vector2 {
	length := math.Sqrt(vec.MagSqr())
	if length != 0 {
		return Vector2{
			X: vec.X / length,
			Y: vec.Y / length,
		}
	}
	return vec
}
func (vec Vector2) UnitVector() Vector2 {
	result := vec

	length := vec.Mag()
	if length != 0 {
		result.X /= length
		result.Y /= length
	}

	return result
}

func (vec Vector2) Normal() Vector2 {
	result := &Vector2{
		X: vec.Y,
		Y: -vec.X,
	}

	return result.Normalize()
}

func (vec Vector2) Dot(vec2 Vector2) float64 {
	return vec.X*vec2.X + vec.Y*vec2.Y
}

func (vec Vector2) Cross(vec2 Vector2) float64 {
	return vec.X*vec2.Y - vec.Y*vec2.X
}

func (vec Vector2) Equals(vec2 Vector2) bool {
	return vec.X == vec2.X && vec.Y == vec2.Y
}

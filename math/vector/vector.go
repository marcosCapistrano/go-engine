package vector

type Vector struct {
	X float32
	Y float32
}

func New(x, y float32) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (vec1 *Vector) Add(vec2 Vector) {
	vec1.X += vec2.X
	vec1.Y += vec2.Y
}

func (vec1 *Vector) Sub(vec2 Vector) {
	vec1.X -= vec2.X
	vec1.Y -= vec2.Y
}

func (vec Vector) Mult(n float32) Vector {
	vec.X *= n
	vec.Y *= n

	return vec
}

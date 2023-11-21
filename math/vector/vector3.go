package vector

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func Cross(vec, vec2 *Vector3) Vector3 {
	result := Vector3{}
	result.X = vec.Y*vec2.Z - vec.Z*vec2.Y
	result.Y = vec.Z*vec2.X - vec.X*vec2.Z
	result.Z = vec.X*vec2.Y - vec.Y*vec2.X

	return result
}

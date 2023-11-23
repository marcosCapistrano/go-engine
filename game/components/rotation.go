package components

type Rotation struct {
	ID    string
	Value float64
}

func (r *Rotation) Mask() uint64 {
	return MaskRotation
}

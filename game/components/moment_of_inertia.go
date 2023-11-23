package components

type MomentOfInertia struct {
	ID      string
	Value   float64
	Inverse float64
}

func (m *MomentOfInertia) Mask() uint64 {
	return MaskMomentOfInertia
}

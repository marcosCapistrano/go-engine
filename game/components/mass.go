package components

type Mass struct {
	ID      string
	Value   float64
	Inverse float64
}

func (m *Mass) Mask() uint64 {
	return MaskMass
}

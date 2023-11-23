package components

type AngularMotion struct {
	ID           string `json:"id"`
	Velocity     float64
	Acceleration float64
	Torque       float64
}

func (am *AngularMotion) Mask() uint64 {
	return MaskAngularMotion
}

// Integrate causes the integration between acceleration, velocity, and position according to deltaTime(dt) and mass
func (am *AngularMotion) Integrate(rotation *float64, invMomentOfInertia float64, dt float64) {
	sumOfTorques := am.Torque * invMomentOfInertia
	am.Acceleration = sumOfTorques
	am.Velocity = am.Velocity + am.Acceleration*dt
	*rotation = *rotation + am.Velocity*dt

	// clear torque
	am.Torque = 0
}

func (am *AngularMotion) AddTorque(torque float64) *AngularMotion {
	am.Torque += torque

	return am
}

/*
func (a *BoxCollider) WithWidth(w int) *BoxCollider {
	a.Width = w
	return a
}

func (a *BoxCollider) WithHeight(h int) *BoxCollider {
	a.Height = h
	return a
}

// NewPosition ...
func NewBoxCollider() ecs.Component {
	return &BoxCollider{
		ID: "box-collider",
	}
}
*/

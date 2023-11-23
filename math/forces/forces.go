package forces

import (
	"engine/math/vector"
	"engine/util"
	"fmt"
)

func Weight(mass float64) vector.Vector2 {
	gravity := vector.Vector2{X: 0, Y: util.GRAVITY_MS * util.PIXELS_PER_METER}
	return gravity.Scale(mass)
}

func Drag(velocity vector.Vector2, k float64) vector.Vector2 {
	dragForce := vector.Vector2{}

	if velocity.MagSqr() > 0 {
		dragDirection := velocity.Normalize().Scale(-1)
		dragMag := k * velocity.MagSqr()

		dragForce = dragDirection.Scale(dragMag)
	}

	return dragForce
}

// vec is velocity
/*
func GenerateFrictionForce(vec Vector2, k float64) Vector2 {
	frictionDir := vec.UnitVector()
	frictionDir.Scale(-1)

	frictionForce := frictionDir.Scale(k)

	return frictionForce
}


*/
// vec is position
func GenerateSpringForce(vec vector.Vector2, anchor vector.Vector2, restLength float64, k float64) vector.Vector2 {
	d := vec.Sub(anchor)

	displacement := d.Mag() - restLength
	fmt.Println("disp", displacement)

	springDirection := d.UnitVector()
	springMag := -k * displacement

	springForce := springDirection.Scale(springMag)

	return springForce
}

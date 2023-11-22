package vector

import "fmt"

// vec is velocity
func GenerateDragForce(vec Vector2, k float64) Vector2 {
	dragForce := Vector2{}

	if vec.MagSqr() > 0 {
		dragDirection := vec.Normalize().Scale(-1)
		dragMag := k * vec.MagSqr()

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
func GenerateSpringForce(vec Vector2, anchor Vector2, restLength float64, k float64) Vector2 {
	d := vec.Sub(anchor)

	displacement := d.Mag() - restLength
	fmt.Println("disp", displacement)

	springDirection := d.UnitVector()
	springMag := -k * displacement

	springForce := springDirection.Scale(springMag)

	return springForce
}

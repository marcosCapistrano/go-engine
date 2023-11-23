package components

import (
	"engine/game/components/base"
	"engine/game/components/base/angular"
	"engine/game/components/base/linear"
)

type RigidBody struct {
	ID string `json:"id"`
	linear.LinearMotion
	angular.AngularMotion
	base.Mass
	base.InvMass
	base.MomentOfInertia
	base.InvMomentOfInertia

	SumOfForces  base.Force
	sumOfTorques base.Torque
}

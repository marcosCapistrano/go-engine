package systems

import (
	"engine/ecs"
	"engine/game/components"
	"fmt"
)

// Movement ...
type Collision struct {
	err       error
	deltaTime *float32
}

func (a *Collision) Error() (err error) {
	return a.err
}

func (a *Collision) Setup() {}

func (a *Collision) Process(registry ecs.Registry) {
	entities := registry.FilterByMask(components.MaskPosition | components.MaskBoxCollider)
	if len(entities) > 1 {
		for i := 0; i < len(entities)-1; i++ {
			entity := entities[i]
			nextEntity := entities[i+1]

			p1 := entity.Get(components.MaskPosition).(*components.Position)
			c1 := entity.Get(components.MaskBoxCollider).(*components.BoxCollider)

			p2 := nextEntity.Get(components.MaskPosition).(*components.Position)
			c2 := nextEntity.Get(components.MaskBoxCollider).(*components.BoxCollider)

			if CheckAABBCollision(int(p1.X), int(p1.Y), c1.Width, c1.Height, int(p2.X), int(p2.Y), c2.Width, c2.Height) {
				fmt.Println("colisaooo")
			}

		}
	}
}

func CheckAABBCollision(aX, aY, aW, aH, bX, bY, bW, bH int) bool {
	return aX < bX+bW && aX+aW > bX && aY < bY+bH && aY+aH > bY
}

func (a *Collision) Teardown() {

}

func (m *Collision) WithData(deltaTime *float32) *Collision {
	m.deltaTime = deltaTime
	return m
}

// NewMovement ...
func NewCollision() ecs.System {
	return &Collision{}
}

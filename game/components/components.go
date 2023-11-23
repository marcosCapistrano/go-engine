package components

const (
	MaskPosition            = uint64(1 << 0)
	MaskSize                = uint64(1 << 1)
	MaskVelocity            = uint64(1 << 2)
	MaskAcceleration        = uint64(1 << 3)
	MaskBoxCollider         = uint64(1 << 4)
	MaskMass                = uint64(1 << 5)
	MaskRotation            = uint64(1 << 6)
	MaskShape               = uint64(1 << 7)
	MaskCenterOfMass        = uint64(1 << 8)
	MaskAngularAcceleration = uint64(1 << 9)
	MaskAngularVelocity     = uint64(1 << 10)
	MaskMomentOfInertia     = uint64(1 << 11)
	MaskRigidBody           = uint64(1 << 12)
	MaskLinearMotion        = uint64(1 << 13)
	MaskAngularMotion       = uint64(1 << 14)
)

package environment

type Environment interface {
	GetState() *State
	WithinTrackBounds() bool
	WithinAngleBounds() bool
	PerformAction(float64)
	Reset()
}

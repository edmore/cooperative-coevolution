package environment

type Environment interface {
	GetState() *State
}

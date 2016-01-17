package environment

type Environment interface {
	GetWorld() *Gridworld
	GetState() *State
	Caught() bool
	Surrounded() bool
	PerformPreyAction()
	PerformPredatorAction()
	Reset()
}

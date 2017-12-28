package environment

type Environment interface {
	GetWorld() *Gridworld
	GetState() *State
	SetPreyPosition(int, int)
	Caught() bool
	Surrounded() bool
	PerformPreyAction(int)
	PerformPredatorAction(int, []float64)
	Reset(int)
}

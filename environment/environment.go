package environment

type Environment interface {
	GetState() *State
	Caught() bool
	Surrounded() bool
	PerformPreyAction()
	PerformPredatorAction()
	Reset()
}

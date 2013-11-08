/*
 Package environment implements the simulation environment(s)
 Current implementation - Pole balancing / Inverted Pendulum Task (double pole balancing with full state information)
*/

package cartpole

var (
	bias             float32 = 0.5
	oneDegree        float64 = 0.0174532 // 2pi/360
	sixDegrees       float64 = 0.1047192
	twelveDegrees    float64 = 0.2094384
	fifteenDegrees   float64 = 0.2617993
	thirtysixDegrees float64 = 0.628329
	sixtyfourDegrees float64 = 1.2566580
	fiftyDegrees     float64 = 0.87266
	inputDimension   int     = 7
	outputDimension  int     = 1
	tau              float64 = 0.01 //seconds between state updates (the time step)
	euler_tau        float64 = tau / 4
	maxFitness       int     = 1000
	poleInc          float64 = 0.05
	massInc          float64 = 0.01
	minInc           float64 = 0.001
	gravity          float64 = 9.8
	RK4              bool    = true
	state                    = make([]float64, inputDimension-1)
	input                    = make([]float64, inputDimension)
	output                   = make([]float64, outputDimension)
)

type Cartpole struct {
	Name       string
	NumOfPoles int
	Markov     bool
	TrackSize  float64
	MUp        float64 // Pole-hinge Friction Coefficient
	MUc        float64 // Cart-track Friction Coefficient
	MassCart   float64
	MassPole1  float64
	MassPole2  float64
	Length1    float64 // actually half the pole's length
	Length2    float64
	ForceMag   float64 // force magnitude in Newtons
}

func NewCartpole(markov bool, numofpoles int) *Cartpole {
	return &Cartpole{
		Name:       "Pole Balancing Task",
		NumOfPoles: numofpoles,
		Markov:     markov,
		TrackSize:  2.4,
		MUp:        0.000002,
		MUc:        0.0005,
		MassCart:   1.0,
		MassPole1:  0.1,
		MassPole2:  0.01,
		Length1:    0.5,
		Length2:    0.05,
		ForceMag:   10.0}
}

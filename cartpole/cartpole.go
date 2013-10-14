/*
Package cartpole implements the Pole balancing / Inverted Pendulum Task
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
)

type Cartpole struct {
	Name       string
	NumOfPoles int
	Markov     boolean
	TrackSize  float64
	MUP        float64
	MUC        float64
	Gravity    float64
	MassCart   float64
	MassPole1  float64
	MassPole2  float64
	Length1    float64 // actually half the pole's length
	Length2    float64
	ForceMag   float64
	TAU        float64 //seconds between state updates

}

func NewCartpole(markov boolean, numofpoles int) *Cartpole {
	return &Cartpole{
		Name:       "Pole Balancing Task",
		NumOfPoles: numofpoles,
		Markov:     markov,
		TrackSize:  2.4,
		MUP:        0.000002,
		MUC:        0.0005,
		Gravity:    -9.8,
		MassCart:   1.0,
		MassPole1:  0.1,
		MassPole2:  0.01,
		Length1:    0.5,
		Length2:    0.05,
		ForceMag:   10.0,
		TAU:        0.01}
}

package network

import (
	"fmt"
	//	"github.com/edmore/esp/population"
	"testing"
)

func TestNewFeedFoward(t *testing.T) {
	const in, hid, out = 3, 2, 1
	// No bias
	ff1 := NewFeedForward(in, hid, out, 0)
	fmt.Println(ff1)
	if x := ff1.GeneSize; x != in+out {
		t.Errorf("Without Bias : GeneSize = %v, we want %v", x, in+out)
	}
	// With bias
	ff2 := NewFeedForward(in, hid, out, 1)
	fmt.Println(ff2)
	if x := ff2.GeneSize; x != in+out+1 {
		t.Errorf("With Bias : GeneSize = %v, we want %v", x, in+out+1)
	}
}

func TestActivate(t *testing.T) {
}

func TestCreate(t *testing.T) {
}

func TestGetTotalInputs(t *testing.T) {
	const in, hid, out = 3, 2, 1
	// With bias
	ff3 := NewFeedForward(in, hid, out, 1)
	if x := ff3.GetTotalInputs(); x != in+1 {
		t.Errorf("With Bias : Total Inputs = %v, we want %v", x, in+1)
	}
	//Without bias
	ff4 := NewFeedForward(in, hid, out, 0)
	if x := ff4.GetTotalInputs(); x != in {
		t.Errorf("Without Bias : Total Inputs = %v, we want %v", x, in)
	}

}

func TestHasBias(t *testing.T) {
	const in, hid, out = 3, 2, 1
	ff5 := NewFeedForward(in, hid, out, 1)
	ff6 := NewFeedForward(in, hid, out, 0)

	if x := ff5.HasBias(); x != true {
		t.Errorf("With Bias : Has bias = %v, we want %v", x, true)
	}

	if x := ff6.HasBias(); x != false {
		t.Errorf("Without Bias : Has Bias = %v, we want %v", x, false)
	}
}

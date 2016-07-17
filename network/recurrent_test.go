package network

import (
	"fmt"
	"testing"
)

func TestNewRecurrent(t *testing.T) {
	const in, hid, out = 3, 2, 1
	// No bias
	rc1 := NewRecurrent(in, hid, out, false)
	fmt.Println(rc1)
    // for a recurrent network the Genesize is "in + hid" vs "in + out"
	if x := rc1.GeneSize; x != in+hid {
		t.Errorf("Without Bias : GeneSize = %v, we want %v", x, in+out)
	}
	// With bias
	rc2 := NewRecurrent(in, hid, out, true)
	fmt.Println(rc2)
	if x := rc2.GeneSize; x != in+hid+1 {
		t.Errorf("With Bias : GeneSize = %v, we want %v", x, in+out+1)
	}
}


func TestRecurrentGetTotalInputs(t *testing.T) {
	const in, hid, out = 3, 2, 1
	// With bias
	rc3 := NewRecurrent(in, hid, out, true)
	if x := rc3.GetTotalInputs(); x != in+1 {
		t.Errorf("With Bias : Total Inputs = %v, we want %v", x, in+1)
	}
	//Without bias
	rc4 := NewRecurrent(in, hid, out, false)
	if x := rc4.GetTotalInputs(); x != in {
		t.Errorf("Without Bias : Total Inputs = %v, we want %v", x, in)
	}

}

func TestRecurrentHasBias(t *testing.T) {
	const in, hid, out = 3, 2, 1
	rc5 := NewRecurrent(in, hid, out, true)
	rc6 := NewRecurrent(in, hid, out, false)

	if x := rc5.HasBias(); x != true {
		t.Errorf("With Bias : Has bias = %v, we want %v", x, true)
	}

	if x := rc6.HasBias(); x != false {
		t.Errorf("Without Bias : Has Bias = %v, we want %v", x, false)
	}
}

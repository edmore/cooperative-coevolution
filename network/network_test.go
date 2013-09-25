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

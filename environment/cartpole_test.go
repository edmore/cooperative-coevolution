package environment

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCartpole(t *testing.T) {
	//c := NewCartpole()
}

func TestReset(t *testing.T) {
	c := NewCartpole(.05)
	c.Reset()
	//	fmt.Println(s)
}

func TestPerformAction(t *testing.T) {
	c := NewCartpole(.05)
	c.Reset()

	fmt.Println("[Initial State:]\n", c.GetState())
	for c.WithinTrackBounds() && c.WithinAngleBounds() {
		c.PerformAction(0)
		elem := reflect.ValueOf(c.GetState()).Elem()
		typeOfCartpole := elem.Type()
		fmt.Println("\n[Update ...]")
		for i := 0; i < elem.NumField(); i++ {
			f := elem.Field(i)
			fmt.Printf("%s = %v | ",
				typeOfCartpole.Field(i).Name, f.Interface())
		}
	}
}

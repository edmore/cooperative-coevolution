package cartpole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCartpole(t *testing.T) {
	//c := NewCartpole()
}

func TestReset(t *testing.T) {
	c := NewCartpole()
	c.Reset()
	//	fmt.Println(s)
}

func TestPerformAction(t *testing.T) {
	c := NewCartpole()
	c.Reset()

	fmt.Println("[Initial State:]\n", c.GetState())
	for c.WithinTrackBounds() && c.WithinAngleBounds() {
		s := c.PerformAction(1)
		elem := reflect.ValueOf(s).Elem()
		typeOfCartpole := elem.Type()
		fmt.Println("\n[Update ...]")
		for i := 0; i < elem.NumField(); i++ {
			f := elem.Field(i)
			fmt.Printf("%s = %v | ",
				typeOfCartpole.Field(i).Name, f.Interface())
		}

	}
	fmt.Println("\n")
}

package util

import (
	"fmt"
	"testing"
)

func TestRoundDown(t *testing.T) {

	var f =0.12345
	fmt.Printf("%f\n", RoundDown(f, 4))
	fmt.Printf("%f\n", Round(f, 4))

}


func TestRandom(t *testing.T) {

	fmt.Printf("%d\n", RandomInt(0, 10))
	fmt.Printf("%d\n", RandomInt(0, 10))
	fmt.Printf("%d\n", RandomInt(0, 10))
	fmt.Printf("%d\n", RandomInt(0, 4))

}
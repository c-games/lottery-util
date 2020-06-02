package util

import (
	"fmt"
	"testing"
)

func TestGetV4(t *testing.T) {

	u, err := GetV4()

	if err != nil {
		t.Fatal(err)

	}
	fmt.Printf("%s", u)

}

func TestGetV4Trim(t *testing.T) {

	u, err := GetV4Trim()

	if err != nil {
		t.Fatal(err)

	}
	fmt.Printf("%s", u)

}

func TestGetV1(t *testing.T) {

	u, err := GetV1()

	if err != nil {
		t.Fatal(err)

	}
	fmt.Printf("%s", u)

}

func TestGetV1Trim(t *testing.T) {

	u, err := GetV1Trim()

	if err != nil {
		t.Fatal(err)

	}
	fmt.Printf("%s", u)

}

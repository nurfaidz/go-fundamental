package main

import (
	"fmt"
	"testing"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Fail()
	}

	fmt.Println("TestFailSum stop execution")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.FailNow()
	}

	fmt.Println("TestSum stop execution")
}

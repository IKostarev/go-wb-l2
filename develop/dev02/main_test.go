package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	strings := "a1b2c3d4"
	expected := "abbcccdddd"

	ok := Unpack(strings)

	if ok == expected {
		fmt.Printf("Тест пройден %v = %v\n", ok, strings)
	} else {
		t.Errorf("Тест не пройден %v != %v\n", ok, strings)
	}
}

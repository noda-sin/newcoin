package main

import (
	"fmt"
	"testing"
)

func TestGetNewCoins(t *testing.T) {
	newCoins, err := getNewCoins()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if len(newCoins) == 0 {
		t.Fatal("failed test")
	}

	fmt.Println(newCoins)
}

package main

import "testing"

func TestContains(t *testing.T) {
	list := []string{"a", "b", "d"}
	if contains(list, "a") != true {
		t.Fatalf("failed test")
	}
	if contains(list, "c") == true {
		t.Fatalf("failed test")
	}
}

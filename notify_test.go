package main

import (
	"testing"
)

func TestNotifyNewCoin(t *testing.T) {
	err := notifyNewCoin()
	if err != nil {
		t.Fatalf("failed test%#v", err)
	}
}

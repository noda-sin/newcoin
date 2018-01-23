package main

import (
	"testing"
)

func TestPostTweet(t *testing.T) {
	err := postTweet("test")
	if err != nil {
		t.Fatalf("failed test%#v", err)
	}
}

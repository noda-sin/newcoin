package main

import (
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	wc := make([]string, 0)
	wc = append(wc, "a")
	wc = append(wc, "b")
	if len(wc) != 2 {
		t.Fatalf("failed test")
	}

	fp, _ := ioutil.TempFile("", "test")
	fpath := fp.Name()
	defer fp.Close()

	err := writeOldCoins(fpath, wc)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	rc, err := readOldCoins(fpath)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(wc) != len(rc) || len(rc) != 2 {
		t.Fatalf("failed test")
	}
}

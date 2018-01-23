package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func writeOldCoins(path string, oldCoins []string) (err error) {
	bytes, err := json.Marshal(oldCoins)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, bytes, os.ModePerm)
	return
}

func readOldCoins(path string) (oldCoins []string, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &oldCoins)
	return
}

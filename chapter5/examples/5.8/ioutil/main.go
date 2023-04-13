package main

import (
	"os"
	"sync"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}

var mu sync.Mutex
var m = make(map[string]int)

func loockup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

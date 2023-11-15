package main

import (
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, World!"

	wg.Add(1)
	var m sync.Mutex
	go updateMessage("Goodbye, cruel world!", &m)
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("incorrect value in msg")
	}
}

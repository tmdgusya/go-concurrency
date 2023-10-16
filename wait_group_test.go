package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	// I always hand pointer to printsomething function
	printSomething("test", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	aa := string(result)

	os.Stdout = stdOut

	fmt.Println(aa)

	if !strings.Contains(aa, "test") {
		t.Errorf("Expected to find test, but it isn't there")
	}
}

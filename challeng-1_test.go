package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_go_routine(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	main()

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	str := string(result)

	os.Stdout = stdOut

	if !strings.Contains(str, "Hello, Universe\nHello, Cosmos\nHello, world!") {
		fmt.Println(str)
		t.Errorf("")
	}
}

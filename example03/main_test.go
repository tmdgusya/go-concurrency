package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// go test -race .
func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	// Close the pipe
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut // Change to original value after calculation

	if !strings.Contains(output, "$34320.00") {
		t.Error("Error!")
	}
}

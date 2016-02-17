package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestWriteLine(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + "\\outFile.test"
	cases := []struct {
		line                     string
		startIndex, outputLength int
		expected                 string
	}{
		{"Write all of me", 0, 0, "Write all of me"},
		{"Write this", 6, 0, "this"},
		{"This will be written", 0, 4, "This"},
		{"Foolish sucka", 4, 8, "ish suck"},
	}
	file, err := os.Create(path)
	for _, c := range cases {
		if err != nil {
			t.Errorf("Encountered an error while testing WriteLine() for path %q. Error: %q", path, err)
		}
		WriteLine(c.line, file, c.startIndex, c.outputLength)
	}
	file.Close()
	file, err = os.Open(path)
	scanner := bufio.NewScanner(file)
	for _, c := range cases {
		if scanner.Scan() != true {
			t.Errorf("WriteLine(%q, %q, %d, %d) did not write to the file, expected %q", c.line, path, c.startIndex, c.outputLength, c.expected)
			file.Close()
			continue
		}
		line := strings.TrimSpace(scanner.Text())
		if strings.Compare(c.expected, line) != 0 {
			t.Errorf("WriteLine(%q, %q, %d, %d) == '%q', expected %q", c.line, path, c.startIndex, c.outputLength, line, c.expected)
		}
	}
	file.Close()
	// Cleanup
	os.Remove(path)
}

func BenchmarkWriteLine(b *testing.B) {
	// Perform setup
	line := "This is a test line to perform a fancy shmancy benchmark test"
	startIndex, outputLength := 0, 0
	path := "outFile.test"
	file, err := os.Create(path)
	if err != nil {
		b.Errorf("Encountered an error while benchmarking WriteLine() for path %q. Error: %q", path, err)
	}
	// Reset the timer
	b.ResetTimer()
	// Run benchmark
	for i := 0; i < b.N; i++ {
		WriteLine(line, file, startIndex, outputLength)
	}
}

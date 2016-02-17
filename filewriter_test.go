package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestWriteLine(t *testing.T) {
	path := "outFile.test"
	cases := []struct {
		line string
		startIndex, outputLength int
		expected string
	} {
		{"Write all of me", 0, 0, "Write all of me"},
		{"Write this", 6, 0, "this"},
		{"This will be written", 0, 4, "This"},
		{"Foolish sucka", 4, 8, "ish suck"},
	}
	for _, c := range cases {
		// Create will truncate the file if it exists.
		file, err := os.Create(path)
		if err != nil {
			t.Errorf("Encountered an error while testing WriteLine() for path %q. Error: %q", path, err)
		}

		WriteLine(c.line, path, c.startIndex, c.outputLength)
		
		scanner := bufio.NewScanner(file)
		// We only care about the first line.
		if scanner.Scan() != true {
			t.Errorf("WriteLine(%q, %q, %d, %d) did not write to the file, expected %q", c.line, path, c.startIndex, c.outputLength, c.expected)
			file.Close()
			continue
		}
		line := scanner.Text()
		if strings.Compare(c.expected, strings.TrimSpace(line)) != 0 {
			t.Errorf("WriteLine(%q, %q, %d, %d) == '%q', expected %q", c.line, path, c.startIndex, c.outputLength, line, c.expected)
		}
		file.Close()
	}
	// Cleanup
	os.Remove(path)
}
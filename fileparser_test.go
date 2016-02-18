package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {
	numberOfLines := 100
	// Set up a file
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("Encountered an error while getting the current directory to test ParseFile(): %q", err)
	}
	path := pwd + "\\TestParseFile.in"
	outPath := pwd + "\\TestParseFile.out"
	file, err := os.Create(path)
	if err != nil {
		t.Errorf("Encountered an error while creating file to test ParseFile() for path %q. Error: %q", path, err)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < numberOfLines; i++ {
		if i%2 == 0 {
			fmt.Fprintln(writer, "Talk about beating a dead horse.")
		} else {
			fmt.Fprintln(writer, "I feel sort of loopy...")
		}
		writer.Flush()
	}
	file.Close()
	filters := map[string][]int{"dead": []int{21, 50}}
	ParseFile(path, outPath, filters, 0, 0)
	file, err = os.Open(outPath)
	if err != nil {
		t.Errorf("Encountered an error while opening file to test ParseFile() for path %q. Error: %q", outPath, err)
	}
	scanner := bufio.NewScanner(file)
	outputLineCount := 0
	for scanner.Scan() {
		outputLineCount++
	}
	if outputLineCount != (numberOfLines / 2) {
		t.Errorf("Expected %d lines in output file, but found %d lines.", (numberOfLines / 2), outputLineCount)
	}
	file.Close()
	os.Remove(path)
	os.Remove(outPath)
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		line     string
		filters  map[string][]int
		expected bool
	}{
		// Test edges of startIndex and length filter values
		{"123456789", map[string][]int{"12": []int{0, 10}}, true},
		{"123456789", map[string][]int{"12": []int{0, 9}}, true},
		{"123456789", map[string][]int{"1": []int{0, 2}}, true},
		{"123456789", map[string][]int{"2": []int{0, 2}}, true},
		{"123456789", map[string][]int{"2": []int{0}}, true},
		// Test line shorter than startIndex
		{"abcde", map[string][]int{"abc": []int{9000, 2}}, false},
		// Test empty filter value
		{"123456789", map[string][]int{"": []int{0, 1}}, true},
		// Test unmatched filter value
		{"abcdefg1234567", map[string][]int{"value": []int{0, 1}}, false},
		// Test filter with no endIndex value
		{"SearchEverythingFromStartIndexOn", map[string][]int{"Everything": []int{0}}, true},
		{"SearchEverythingFromStartIndexOn", map[string][]int{"Everything": []int{6}}, true},
		{"SearchEverythingFromStartIndexOn", map[string][]int{"Everything": []int{15}}, false},
	}
	for _, c := range cases {
		got := ParseLine(c.line, c.filters)
		if got != c.expected {
			t.Errorf("ReadLine(%q, %+v) == %t, expected %t", c.line, c.filters, got, c.expected)
		}
	}
}

func BenchmarkParseLine(b *testing.B) {
	// Perform setup
	line := "SearchTheWholeString"
	filter := map[string][]int{"Str": []int{0}}
	// Reset the timer
	b.ResetTimer()
	// Run benchmark
	for i := 0; i < b.N; i++ {
		ParseLine(line, filter)
	}
}

func BenchmarkParseLineParallel(b *testing.B) {
	// Perform setup
	line := "SearchTheWholeString"
	filter := map[string][]int{"Str": []int{0}}
	// Reset the timer
	b.ResetTimer()
	// Run benchmark
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ParseLine(line, filter)
		}
	})
}

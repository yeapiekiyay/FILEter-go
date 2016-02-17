package main

import "testing"

func TestParseLine(t *testing.T) {
	cases := []struct {
		line string
		filters map[string][]int
		expected bool
	} {
		// Test edges of startIndex and endIndex filter values
		{"123456789", map[string][]int { "12": []int{0, 1}}, true},
		{"123456789", map[string][]int { "1": []int{0, 1}}, true},
		{"123456789", map[string][]int { "2": []int{0, 1}}, true},
		// Test empty filter value
		{"123456789", map[string][]int { "": []int{0, 1}}, false},
		// Test unmatched filter value
		{"abcdefg1234567", map[string][]int { "value": []int{0, 1}}, false},
		// Test filter with no endIndex value 
		{"SearchEverythingFromStartIndexOn", map[string][]int {"Everything": []int{0}}, true},
		{"SearchEverythingFromStartIndexOn", map[string][]int {"Everything": []int{6}}, true},
		{"SearchEverythingFromStartIndexOn", map[string][]int {"Everything": []int{15}}, false},
	}
	for _, c := range cases {
		got := ParseLine(c.line, c.filters)
		if got != c.expected {
			t.Errorf("ReadLine(%q, %q) == %q, expected %q", c.line, c.filters, got, c.expected)
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
	for i:= 0; i < b.N; i++ {
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
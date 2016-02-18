package main

import (
	"testing"
)

func TestNewArguments(t *testing.T) {
	cases := []struct {
		inputFile, outputFile          string
		outputStartIndex, outputLength int
		filters                        []string
	}{
		{"test.in", "test.out", 0, 0, []string{"value,0,0"}},
		{"test.in", "", 0, 0, []string{"wheeeee,5,2"}},
		{"yo", "dawg", 90, 0, []string{"blah,0,0"}},
		{"test.in", "test.out", 0, 38, []string{"value,0,0"}},
		{"hey", "boo", 0, 0, []string{"ooo,0,0"}},
	}
	for _, c := range cases {
		got := NewArguments(c.inputFile, c.outputFile, c.outputStartIndex, c.outputLength, c.filters)
		if got.InputFile != c.inputFile ||
			got.OutputFile != c.outputFile ||
			got.OutputStartIndex != c.outputStartIndex ||
			got.OutputLength != c.outputLength ||
			!CompareStringArrays(got.Filters, c.filters) {
			t.Errorf("Error testing NewArguments. Expected %+v, but got %+v", c, got)
		}
	}
}

func TestMain(t *testing.T) {
	t.Skip("Not testing main since no functionality resides here. All functional components will be unit tested.")
}

func TestValidateArguments(t *testing.T) {
	cases := []struct {
		args     Arguments
		expected bool
	}{
		{*NewArguments("inFile.test", "outFile.test", 0, 0, []string{"A,0,0"}), true},
	}
	for _, c := range cases {
		got, _ := ValidateArguments(c.args)
		if got != c.expected {
			t.Errorf("Error testing ValidateArguments for Arguments %q. Expected %t, got %t", c.args, c.expected, got)
		}
	}
}

func TestParseFilters(t *testing.T) {
	cases := []struct {
		filters  []string
		expected map[string][]int
	}{
		{[]string{"blah,0,0"}, map[string][]int{"blah": []int{0, 0}}},
		{[]string{"whatever,96,0"}, map[string][]int{"whatever": []int{96, 0}}},
		{[]string{"wellthen,0,80"}, map[string][]int{"wellthen": []int{0, 80}}},
		{[]string{"yodawgiheardyoulikeunittests,5,7"}, map[string][]int{"yodawgiheardyoulikeunittests": []int{5, 7}}},
		{[]string{"blah,0"}, map[string][]int{"blah": []int{0, 0}}},
	}
	for _, c := range cases {
		got := ParseFilters(c.filters)
		if !CompareFilters(got, c.expected) {
			t.Errorf("Error testing ParseFilters. Expected %+v, but got %+v", c.expected, got)
		}
	}
}

func CompareStringArrays(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareIntArrays(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareFilters(a, b map[string][]int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if _, ok := b[k]; ok {
			if !CompareIntArrays(v, b[k]) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

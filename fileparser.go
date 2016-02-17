package main

import (
	"strings"
	"unicode/utf8"
)

func ParseFile(fileName string) {

}

func ParseLine(line string, filters map[string][]int) (exportLine bool) {
	filterMatches := make(map[int]bool)
	// This could be more efficient if we mapped unique starting indices to filters and the optional lengths.
	// Then iterated through the list of filters for each unique starting index.
	// If we get a match, skip the rest of the filters for this starting index.
	for k, v := range filters {
		filter := k
		lineLength := utf8.RuneCountInString(line)
		startIndex := v[0]
		searchString := line
		// If we have a length of > 1, we have a length argument for the filter.
		if len(v) > 1 && v[1] > 0 && v[1] <= lineLength {
			length := v[1]
			searchString = line[startIndex:length]
		} else {
			searchString = line[startIndex:lineLength]
		}
		// Set the filterMatches value to true if we already had a match or if we just had one.
		filterMatches[startIndex] = filterMatches[startIndex] || strings.Contains(searchString, filter)
	}
	// Check if we didn't match any filters for each unique starting index.
	for _, v := range filterMatches {
		if v == false {
			return false
		}
	}
	// If we get here, we matched a filter for each unique starting index.
	return true
}

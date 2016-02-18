package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseFile(inFilePath, outFilePath string, filters map[string][]int, outputStartIndex, outputLength int) {
	inFile, err := os.Open(inFilePath)
	defer inFile.Close()
	if err != nil {
		panic(fmt.Sprintf("The input file could not be opened. Input file: %q", inFilePath))
	}
	outFile, err := os.Create(outFilePath)
	defer outFile.Close()
	if err != nil {
		panic(fmt.Sprintf("The output file could not be created. Output file: %q", outFilePath))
	}
	scanner := bufio.NewScanner(inFile)
	// var writeLock = make(chan int, 1)
	for scanner.Scan() {
		text := scanner.Text()
		// Works synchronously...
		writeLine := ParseLine(text, filters)
		if writeLine {
			WriteLine(text, outFile, outputStartIndex, outputLength)
		}
		// go func(line string, filters map[string][]int) {
		// 	writeLine := ParseLine(line, filters)
		// 	if writeLine {
		// 		counter++
		// 		fmt.Sprintln(counter)
		// 		// // Lock access to write to the file
		// 		writeLock <- 1
		// 		WriteLine(line, outFile, outputStartIndex, outputLength)
		// 		<-writeLock
		// 	}
		// }(text, filters)
	}
}

func ParseLine(line string, filters map[string][]int) (exportLine bool) {
	filterMatches := make(map[int]bool)
	// This could be more efficient if we mapped unique starting indices to filters and the optional lengths.
	// Then iterated through the list of filters for each unique starting index.
	// If we get a match, skip the rest of the filters for this starting index.
	for k, v := range filters {
		filter := k
		lineLength := len(line)
		startIndex := v[0]
		searchString := line
		if startIndex > lineLength {
			continue
		}
		// If we have a length of > 1, we have a length argument for the filter.
		if len(v) > 1 && v[1] > 0 && v[1] <= lineLength {
			endIndex := startIndex + v[1]
			searchString = line[startIndex:endIndex]
		} else {
			searchString = line[startIndex:lineLength]
		}
		// Set the filterMatches value to true if we already had a match or if we just had one.
		filterMatches[startIndex] = filterMatches[startIndex] || strings.Contains(searchString, filter)
	}
	// Check if we didn't match any filters for each unique starting index.
	if len(filterMatches) == 0 {
		return false
	}
	for _, v := range filterMatches {
		if v == false {
			return false
		}
	}
	// If we get here, we matched a filter for each unique starting index.
	return true
}

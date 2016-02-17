package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `FILEter - A flat file filtering utility.

	Usage:
	  FILEter -i <inputFile> [-o <outputFile>] [-s <startIndex>] [-l <length>] [--] <filter>...
	  FILEter --inputFile <inputFile> [--outputFile <outputFile>] [--outputStartIndex <startIndex>] [--outputLength <length>] [--] <filter>...	  
	  FILEter (-h | --help)
	  FILEter (-v | --version)

	Options:
	  -i --inputFile  		The file to be filtered.
	  -o --outputFile	  	Optional. The file to write matching results to.
	  -s --outputStartIndex  	Optional. The index in each line that passes the filters at which to start writing to the output file.
	  -l --outputLength  	Optional. The length of the output from the startIndex in each line that passes the filters to write to the output file.
	  -h --help  		Show this screen.
	  -v --version  		Show version.

	Filter argument explained:
	  The filter argument must be in the following format: value,startIndex[,length]
	  The value is the value to be matched in each line of the file being processed. Each value must be unique.
	  The startIndex is the index of the character in the line to begin searching for the value.
	  The length is the length of the string to search for the value, starting from the startIndex. This value is optional.
	  The line will be exported if for each unique startIndex, one of the corresponding filters matches.`

	arguments, _ := docopt.Parse(usage, nil, true, "FILEter 0.1", false)
	fmt.Println(arguments)
}

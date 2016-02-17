package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `FILEter - A flat file filtering utility.

	Usage:
	  FILEter -i <inputFile> [-o <outputFile>] [-s <startIndex>] [-e <endIndex>] [--] <filter>...
	  FILEter --inputFile <inputFile> [--outputFile <outputFile>] [--outputStartIndex <startIndex>] [--outputEndIndex <endIndex>] [--] <filter>...	  
	  FILEter (-h | --help)
	  FILEter (-v | --version)

	Options:
	  -i --inputFile  		The file to be filtered.
	  -o --outputFile	  	Optional. The file to write matching results to.
	  -s --outputStartIndex  	Optional. The index in each line that passes the filters at which to start writing to the output file.
	  -e --outputEndIndex  	Optional. The index in each line that passes the filters at which to stop writing to the output file.
	  -h --help  		Show this screen.
	  -v --version  		Show version.

	Filter argument explained:
	  The filter argument must be in the following format: value,startIndex[,endIndex]
	  The value is the value to be matched in each line of the file being processed.
	  The startIndex is the index of the character in the line to begin searching for the value.
	  The endIndex is the index of the character in the line to end searching for the value. This value is optional.
	  The line will be exported if for each unique startIndex, one of the corresponding filters matches.`

	  arguments, _ := docopt.Parse(usage, nil, true, "FILEter 0.1", false)
	  fmt.Println(arguments)
}
package main

import (
	"github.com/alexflint/go-arg"
	"strconv"
	"strings"
)

func main() {
	var args struct {
		InputFile        string   `arg:"-i,required,help: The input file to filter."`
		OutputFile       string   `arg:"-o,help: The output file to export lines matching the filters to."`
		OutputStartIndex int      `arg:"-s,help: The index of the character in each line that matches to start exporting at."`
		OutputLength     int      `arg:"-l,help: The number of characters from the outputStartIndex to export in each matching line."`
		Filter           []string `arg:"positional,help: One or more filters separated by spaces in the format "value,startIndex[,length].""`
	}
	args.OutputFile = "filtered.out"
	args.OutputStartIndex = 0
	args.OutputLength = 0
	p := arg.MustParse(&args)

	var filters = make(map[string][]int)
	if args.InputFile == "" {
		p.Fail("You must provide an inputFile.")
	}
	if len(args.Filter) == 0 {
		p.Fail("At least one filter must be provided.")
	}
	for _, f := range args.Filter {
		splitString := strings.Split(f, ",")
		if len(splitString) < 2 {
			p.Fail("Each filter must have a value and a startIndex.")
		}
		if splitString[0] == "" {
			p.Fail("Each filter must have a value and a startIndex.")
		}

		value := splitString[0]
		startIndex64, err := strconv.ParseInt(splitString[1], 10, 32)
		if err != nil {
			panic("An error occurred while trying to parse one of the filter's startIndex values. Please check your arguments and try again.")
		}
		startIndex := int(startIndex64)
		length := int(0)
		if len(splitString) == 3 {
			length64, err := strconv.ParseInt(splitString[2], 10, 32)
			length = int(length64)
			if err != nil {
				panic("An error occurred while trying to parse one of the filter's length values. Please check your arguments and try again.")
			}
		}

		filters[value] = []int{startIndex, length}
	}

	ParseFile(args.InputFile, args.OutputFile, filters, args.OutputStartIndex, args.OutputLength)
}

package main

import (
	"github.com/alexflint/go-arg"
	"strconv"
	"strings"
)

func main() {
	args := ParseArguments()
	filters := ParseFilters(args.Filters)
	ParseFile(args.InputFile, args.OutputFile, filters, args.OutputStartIndex, args.OutputLength)
}

func ParseArguments() Arguments {
	var args Arguments
	args.OutputFile = "filtered.out"
	args.OutputStartIndex = 0
	args.OutputLength = 0

	p := arg.MustParse(&args)

	argumentsAreValid, err := ValidateArguments(args)
	if !argumentsAreValid {
		p.Fail(err)
	}

	return args
}

func ValidateArguments(args Arguments) (bool, string) {
	if args.InputFile == "" {
		return false, "You must provide an inputFile."
	}
	if len(args.Filters) == 0 {
		return false, "At least one filter must be provided."
	}
	for _, f := range args.Filters {
		splitString := strings.Split(f, ",")
		if len(splitString) < 2 {
			return false, "Each filter must have a value and a startIndex."
		}
		if splitString[0] == "" {
			return false, "Each filter must have a value and a startIndex."
		}
	}
	return true, ""
}

func ParseFilters(filters []string) map[string][]int {
	var newFilters = make(map[string][]int)
	for _, f := range filters {
		splitString := strings.Split(f, ",")

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

		newFilters[value] = []int{startIndex, length}
	}

	return newFilters
}

func NewArguments(inputFile, outputFile string, outputStartIndex, outputLength int, filters []string) *Arguments {
	return &Arguments{inputFile, outputFile, outputStartIndex, outputLength, filters}
}

type Arguments struct {
	InputFile        string   `arg:"-i,required,help: The input file to filter."`
	OutputFile       string   `arg:"-o,help: The output file to export lines matching the filters to."`
	OutputStartIndex int      `arg:"-s,help: The index of the character in each line that matches to start exporting at."`
	OutputLength     int      `arg:"-l,help: The number of characters from the outputStartIndex to export in each matching line."`
	Filters          []string `arg:"positional,help: One or more filters separated by spaces in the format value.startIndex[.length] with commas in place of periods."`
}

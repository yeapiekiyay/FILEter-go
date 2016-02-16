package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `FILEter - A flat file filtering utility.

	Usage:
	  FILEter (-i | --inputFile) <inputFile>
	  FILEter (-i | --inputFile) <inputFile> [(-o | --outputFile) <outputFile>]
	  FILEter (-i | --inputFile) <inputFile> [(-o | --outputFile) <outputFile>]
	  FILEter <inputFile>
	  FILEter <inputFile> [<outputFile>]
	  FILEter -h | --help
	  FILEter -v | --version

	Options:
	  -i --inputFile	The file to be filtered.
	  -o -outputFile	Optional. The file to write matching results to.
	  -h --help	Show this screen.
	  -v --version	Show version.`
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func WriteLine(line string, file *os.File, startIndex, outputLength int) error {
	writer := bufio.NewWriter(file)
	lineLength := utf8.RuneCountInString(line)
	endIndexExclusive := startIndex + outputLength
	if endIndexExclusive > lineLength || outputLength == 0 {
		endIndexExclusive = lineLength
	}
	toWrite := line[startIndex:endIndexExclusive]
	fmt.Fprintln(writer, toWrite)
	return writer.Flush()
}

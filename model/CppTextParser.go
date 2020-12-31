// Implements an actual parser for CPP code
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"fmt"
	"strings"
)

type GetKeyError struct {}

type CppTextParser struct {}

// Custom error implementation
func (GetKeyError) Error() string {
	return "GetKeyError"
}

// to implement ITextParser interface
func (*CppTextParser) ParseLine(line string) error {
	// TODO
	fmt.Println("CppTextParser - ParseLine calledi with", line)
	return nil
}

func (*CppTextParser) GetKeyFromLine(line string) (str string, e error) {
	// Let's trim the string first: this contains both a space and a tab to be trimmed out
	trimmed := strings.TrimLeft(line, " 	")
	if len(trimmed) > 0 {
		tokens := strings.Split(trimmed, " ")
		if len(tokens) > 0 {
			str = tokens[0]
			e = nil
			return
		}
	}
	str = ""
	e = GetKeyError{}
	return
}

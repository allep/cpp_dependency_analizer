// Implements an actual parser for CPP code
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"fmt"
)

type CppTextParser struct {
	aField string
}

// to implement ITextParser interface
func (*CppTextParser) ParseLine(line string) error {
	// TODO
	fmt.Println("CppTextParser - ParseLine calledi with", line)
	return nil
}

func (*CppTextParser) GetKeyFromLine(line string) (str string, e error) {
	// TODO
	fmt.Println("CppTextParser - GetKeyFromLine called with", line)
	str = ""
	e = nil
	return
}

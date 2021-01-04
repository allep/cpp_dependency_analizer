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
	
	// logic
	// 1. Get key from line
	// 2. Based on key push the new decoder on the stack
	// 3. Based on the parse result of the decoder, decide if the decoder needs to be
	//    popped or not.


	/* 
	Better description
	- I shouldn't resort on some external logic to get the key etc
	- Once I get inside this, I should:
		1. Take the current top decoder
		2. Pass the present line to it
		3. Return
	- It' the decoder who is responsible to;
		- From current line get the key
		- Based on the key decide if a new decoder has to be put onto the stack
		- Based on current line decide if the current decoder has to be popped out of the stack
	
	- What do I need then?
		1. CppTextParser should have a DecoderStack
		2. Every time we parse a line we should consider the decoder on top of the stack and call ParseLine
		   on it
		3. The decoder's parse line output should let us understand if the decoder has to be popped out of
		   the stack

	*/
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

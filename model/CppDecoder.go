// Implements the Cpp decoder logic
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
	"strings"
)

// types used to implement the IDecoder interface for each statement
type CppDecoderInclude struct {
	symbols []string
}

type CppDecoderTypedef struct {}
type CppDecoderEnum struct {}
type CppDecoderClass struct {}
type CppDecoderSingleLineComment struct {}
type CppDecoderMultiLineComment struct {}

// functions

// Index is zero-based
func GetTokenFromLine(line string, index int) (string, error) {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) > 0 {
		tokens := strings.Split(trimmed, " ")
		if len(tokens) > 0 {
			// now let's try to get the correct index out of the line
			if ((index >= 0) && (index < len(tokens))) {
				token := tokens[index]
				return token, nil
			}
		}
	}
	err := errors.New("Tokenization error")
	return "", err
}


// methods

func (c *CppDecoderInclude) PushBack(symbol string) {
	// do not allow for duplicates
	for _, val := range c.symbols {
		if val == symbol {
			return
		}
	}
	// not found, let's append it
	c.symbols = append(c.symbols, symbol)
}

// Precondition of this function: the caller should have already found that #include is present
// inside "line"
func (c *CppDecoderInclude) DecodeLine(line string) (bool, error) {
	token, err := GetTokenFromLine(line, 1) // zero based
	if err == nil {
		// properly sanitize it, by removing:
		// - parenthesis
		// - double quotes
		// - relative paths
		sanitized := strings.TrimLeft(token, " <\"./")
		sanitized = strings.TrimRight(sanitized, " >\"")

		if len(sanitized) > 0 {
			c.PushBack(sanitized)
			// decoder must be popped out
			return true, nil
		}
	}
	// failure
	ret_err := errors.New("Decode error")
	return true, ret_err
}

func (c *CppDecoderInclude) GetSymbols() []string {
	return c.symbols
}

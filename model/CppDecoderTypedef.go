// Implements the Cpp decoder logic for typedefs
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
)

type CppDecoderTypedef struct {
	symbols []string
}

func (d *CppDecoderTypedef) DecodeLine(line string) (bool, error) {
	// TODO
	return true, errors.New("CppDecoderTypedef DecodeLine error")
}

func (d *CppDecoderTypedef) GetDecoderDescription() string {
	return "CppDecoderTypedef"
}

func (d *CppDecoderTypedef) GetSymbols() []string {
	return d.symbols
}

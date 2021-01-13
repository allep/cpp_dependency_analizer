// Implements the Cpp decoder logic for classes
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
)

type CppDecoderClass struct {
	symbols []string
}

func (d *CppDecoderClass) DecodeLine(line string) (bool, error) {
	// TODO
	return true, errors.New("CppDecoderClass DecodeLine error")
}

func (d *CppDecoderClass) GetDecoderDescription() string {
	return "CppDecoderClass"
}

func (d *CppDecoderClass) GetSymbols() []string {
	return d.symbols
}

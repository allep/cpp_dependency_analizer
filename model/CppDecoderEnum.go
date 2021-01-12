// Implements the Cpp decoder logic for enums 
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
)

type CppDecoderEnum struct {}

func (d *CppDecoderEnum) DecodeLine(line string) (bool, error) {
	// TODO
	return true, errors.New("CppDecoderEnum DecodeLine error")
}

func (d *CppDecoderEnum) GetDecoderDescription() string {
	return "CppDecoderEnum"
}


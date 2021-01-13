// Implements the Cpp decoder logic for typedefs
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
)

type CppDecoderTypedef struct{}

func (d *CppDecoderTypedef) DecodeLine(line string) (bool, error) {
	// TODO
	return true, errors.New("CppDecoderTypedef DecodeLine error")
}

func (d *CppDecoderTypedef) GetDecoderDescription() string {
	return "CppDecoderTypedef"
}

// Implements the Cpp decoder logic for classes
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
)

type CppDecoderClass struct {
	class_observer ClassObserver
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

func (d *CppDecoderClass) Flush() error {
	if d.class_observer == nil {
		return errors.New("Invalid class observer")
	}
	d.class_observer.UpdateClassList(d.symbols)
	return nil
}

func (d *CppDecoderClass) SetClassObserver(obs ClassObserver) {
	d.class_observer = obs
}


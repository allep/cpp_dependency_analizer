// Implements the Cpp decoder logic for enums
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
)

type CppDecoderEnum struct {
	enum_observer EnumObserver
	symbols []string
}

func (d *CppDecoderEnum) DecodeLine(line string) (bool, error) {
	// TODO
	return true, errors.New("CppDecoderEnum DecodeLine error")
}

func (d *CppDecoderEnum) GetDecoderDescription() string {
	return "CppDecoderEnum"
}

func (d *CppDecoderEnum) GetSymbols() []string {
	return d.symbols
}

func (d *CppDecoderEnum) Flush() error {
	if d.enum_observer == nil {
		return errors.New("Invalid enum observer")
	}
	d.enum_observer.UpdateEnumList(d.symbols)
	return nil
}

func (d *CppDecoderEnum) SetEnumObserver(obs EnumObserver) {
	d.enum_observer = obs
}


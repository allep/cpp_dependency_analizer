// Implements the Cpp decoder logic for typedefs
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
)

type CppDecoderTypedef struct {
	typedef_observer TypedefObserver
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

func (d *CppDecoderTypedef) Flush() error {
	if d.typedef_observer == nil {
		return errors.New("Invalid typedef observer")
	}
	d.typedef_observer.UpdateTypedefList(d.symbols)
	return nil
}

func (d *CppDecoderTypedef) SetTypedefObserver(obs TypedefObserver) {
	d.typedef_observer = obs
}


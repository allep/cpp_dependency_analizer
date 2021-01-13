// Implements the CPP decoder state machine
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
	"fmt"
	"github.com/allep/cpp_dependency_analyzer/core"
)

type CppDecoderFSM struct {
	parser *CppTextParser
	decoder_stack []core.IDecoder
}

// methods

func (f *CppDecoderFSM) SetObserver(obs *CppTextParser) {
	f.parser = obs
}

func (f *CppDecoderFSM) Update(key string) (state_changed bool) {
	var p core.IDecoder
	switch key {
	case "#include":
		fmt.Println("Include case")
		pi := new(CppDecoderInclude)
		pi.SetIncludeObserver(f.parser)
		p = pi
	case "class":
		fmt.Println("Class case")
		pc := new(CppDecoderClass)
		pc.SetClassObserver(f.parser)
		p = pc
	case "typedef":
		fmt.Println("Typedef case")
		pt := new(CppDecoderTypedef)
		pt.SetTypedefObserver(f.parser)
		p = pt
	case "enum":
		fmt.Println("Enum case")
		pe := new(CppDecoderEnum)
		pe.SetEnumObserver(f.parser)
		p = pe
	default:
		fmt.Println("Default case")
	}
	if p != nil {
		f.decoder_stack = append(f.decoder_stack, p)
		state_changed = true
	}
	return
}

func (f *CppDecoderFSM) Pop() (bool, error) {
	if len(f.decoder_stack) == 0 {
		return false, errors.New("Pop error")
	}

	f.decoder_stack = f.decoder_stack[:len(f.decoder_stack)-1]
	return true, nil
}

func (f *CppDecoderFSM) StackSize() int {
	return len(f.decoder_stack)
}

func (f *CppDecoderFSM) GetCurrentState() (*core.IDecoder, error) {
	if len(f.decoder_stack) == 0 {
		return nil, errors.New("Empty stack")
	}
	return &f.decoder_stack[len(f.decoder_stack)-1], nil
}

func (f *CppDecoderFSM) GetCurrentStateDescription() (description string, err error) {
	if len(f.decoder_stack) == 0 {
		return "", errors.New("Empty stack")
	}
	return f.decoder_stack[len(f.decoder_stack)-1].GetDecoderDescription(), nil
}


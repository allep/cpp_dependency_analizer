// Implements the CPP decoder state machine
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"errors"
	"fmt"
	"github.com/allep/cpp_dependency_analyzer/core"
)

type CppDecoderFSM struct {
	decoder_stack []core.IDecoder
}

// methods

func (f *CppDecoderFSM) Update(key string) (state_changed bool) {
	var p core.IDecoder
	switch key {
	case "#include":
		fmt.Println("Include case")
		p = new(CppDecoderInclude)
	case "class":
		fmt.Println("Class case")
		p = new(CppDecoderClass)
	case "typedef":
		fmt.Println("Typedef case")
		p = new(CppDecoderTypedef)
	case "enum":
		fmt.Println("Enum case")
		p = new(CppDecoderEnum)
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

func (f *CppDecoderFSM) GetCurrentStateDescription() (description string, err error) {
	if len(f.decoder_stack) == 0 {
		return "", errors.New("Empty stack")
	}
	return f.decoder_stack[len(f.decoder_stack)-1].GetDecoderDescription(), nil
}


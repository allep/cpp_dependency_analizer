// Implements an actual parser for CPP code
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"fmt"
	"errors"
	"strings"
)

type CppTextParser struct {
	include_list []string

	// The decoders' FSM
	fsm CppDecoderFSM
}

func (p *CppTextParser) Init() {
	p.fsm.SetObserver(p)
}

func (p *CppTextParser) GetIncludes() []string {
	return p.include_list
}

// to implement ITextParser interface
func (p *CppTextParser) ParseLine(line string) error {
	// Step 1: try to get a key from current line and evolve internal state machine accordingly.
	key, k_err := p.GetKeyFromLine(line)
	if k_err == nil {
		p.fsm.Update(key)
	}

	// Step 2: use current state to perform current line decoding.
	var state_pop bool
	var dec_err error
	cur_dec, s_err := p.fsm.GetCurrentState()
	if s_err == nil {
		state_pop, dec_err = (*cur_dec).DecodeLine(line)
		if dec_err != nil {
			fmt.Println("CppTextParser: error while decoding line: ", line)
		}
	}
	// Step 3: if we need to pop current decoder, then collect symbols first, then actually pop it.
	if state_pop {
		// TODO
		// This needs to be done transparently. The decoder knows what symbol it produces, so 
		// we can leave the decoder to properly update them, e.g.:
		// cur_dec.Flush()

		// FIXME: this is just a temporary workaround
		f_err := (*cur_dec).Flush()
		if f_err != nil {
			fmt.Println("Flush error:", f_err)
		}
		p.fsm.Pop()
	}

	return nil
}

func (p *CppTextParser) GetKeyFromLine(line string) (str string, e error) {
	// Let's trim the string first: this contains both a space and a tab to be trimmed out
	trimmed := strings.TrimSpace(line)
	if len(trimmed) > 0 {
		tokens := strings.Split(trimmed, " ")
		if len(tokens) > 0 {
			str = tokens[0]
			e = nil
			return
		}
	}
	str = ""
	e = errors.New("Key error")
	return
}

// observer methods

func (p *CppTextParser) UpdateIncludeList(list []string) {
	fmt.Println("CppTextParser: UpdateIncludeList")
	for _, v := range list {
		var found bool = false
		for _, w := range p.include_list {
			if v == w {
				found = true;
				break
			}
		}
		if !found {
			p.include_list = append(p.include_list, v)
		}
	}
}

func (p *CppTextParser) UpdateTypedefList(list []string) {
	fmt.Println("CppTextParser: UpdateTypedefList")
	// TODO
}

func (p *CppTextParser) UpdateEnumList(list []string) {
	fmt.Println("CppTextParser: UpdateEnumList")
	// TODO
}

func (p *CppTextParser) UpdateClassList(list []string) {
	fmt.Println("CppTextParser: UpdateClassList")
	// TODO
}


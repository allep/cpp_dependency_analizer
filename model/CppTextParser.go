// Implements an actual parser for CPP code
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

import (
	"fmt"
	"errors"
	"strings"
)

type CppTextParser struct {
	// FIXME: this is just to test new developments
	symbols []string

	// The decoders' FSM
	fsm CppDecoderFSM
}

func (p *CppTextParser) GetIncludes() []string {
	return p.symbols
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
		cur_state_symbols := (*cur_dec).GetSymbols()
		for _, v := range cur_state_symbols {
			// TODO FIXME: also in case of workaround here we need to avoid duplications
			p.symbols = append(p.symbols, v)
		}
		p.fsm.Pop()
	}

	return nil

	// logic
	// 1. Get key from line
	// 2. Based on key push the new decoder on the stack
	// 3. Based on the parse result of the decoder, decide if the decoder needs to be
	//    popped or not.

	/*
		Better description
		- I shouldn't resort on some external logic to get the key etc
		- Once I get inside this, I should:
			1. Take the current top decoder
			2. Pass the present line to it
			3. Return
		- It' the decoder who is responsible to;
			- From current line get the key
			- Based on the key decide if a new decoder has to be put onto the stack
			- Based on current line decide if the current decoder has to be popped out of the stack

		- What do I need then?
			1. CppTextParser should have a DecoderStack
			2. Every time we parse a line we should consider the decoder on top of the stack and call ParseLine
			   on it
			3. The decoder's parse line output should let us understand if the decoder has to be popped out of
			   the stack

	*/
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

// Implements unit tests for CppDecoderFSM

package model

import (
	"fmt"
	"testing"
)

func TestCppDecoderFSMUpdate(t *testing.T) {
	fmt.Println("Testing CppDecoderFSM > Update")

	var test_cases = []struct {
		key string
		expected_state_change bool
		expected_decoder_desc string
		expected_stack_size int
	}{
		{"invalid", false, "", 0},
		{"#include", true, "CppDecoderInclude", 1},
		{"class", true, "CppDecoderClass", 2},
		{"typedef", true, "CppDecoderTypedef", 3},
		{"enum", true, "CppDecoderEnum", 4},
	}
	fsm := new(CppDecoderFSM)
	
	for _, val := range test_cases {
		changed := fsm.Update(val.key)
		if changed != val.expected_state_change {
			t.Error("FSM state change mismatch.")
		} else {
			t.Log("FSM state changed correctly.")
		}
		desc, _ := fsm.GetCurrentStateDescription()
		if desc != val.expected_decoder_desc {
			t.Error("FSM state decoder description mismatch.")
		} else {
			t.Log("FSM state decoder description matches.")
		}
		stack_size := fsm.StackSize()
		if stack_size != val.expected_stack_size {
			t.Error("FSM state stack size mismatch.")
		} else {
			t.Log("FSM state stack size matches.")
		}
	}
}

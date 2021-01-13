// Implements unit tests for CppDecoderFSM

package model

import (
	"fmt"
	"testing"
)

// Test logic:
// - Call Update() with several possible keys, both valid and invalid.
// - Verify that the corresponding state change happens.
// - Verify the stack is composed of the correct number of stacked decoders.
// - At the end of the test Pop() all decoders, one by one.
func TestCppDecoderFSMUpdate(t *testing.T) {
	var test_cases = []struct {
		key                   string
		expected_state_change bool
		expected_decoder_desc string
		expected_stack_size   int
	}{
		{"invalid", false, "", 0},
		{"#include", true, "CppDecoderInclude", 1},
		{"class", true, "CppDecoderClass", 2},
		{"class", true, "CppDecoderClass", 3}, // simulate an inner class
		{"typedef", true, "CppDecoderTypedef", 4},
		{"invalid", false, "CppDecoderTypedef", 4}, // simulate an invalid state in the between
		{"enum", true, "CppDecoderEnum", 5},
		{"enum", true, "CppDecoderEnum", 6}, // simulate a kind of impossible enum inside another enum
	}
	fsm := new(CppDecoderFSM)

	// Actual loop over possible state changes
	for _, val := range test_cases {
		changed := fsm.Update(val.key)
		if changed != val.expected_state_change {
			t.Error("FSM state change mismatch.")
		} else {
			fmt.Println("FSM state changed correctly.")
		}
		desc, _ := fsm.GetCurrentStateDescription()
		if desc != val.expected_decoder_desc {
			t.Error("FSM state decoder description mismatch.")
		} else {
			fmt.Println("FSM state decoder description matches.")
		}
		stack_size := fsm.StackSize()
		if stack_size != val.expected_stack_size {
			t.Error("FSM state stack size mismatch.")
		} else {
			fmt.Println("FSM state stack size matches.")
		}
	}

	// Now let's pop all decoders from the stack
	stack_size := fsm.StackSize()
	expected_stack_size := stack_size
	for ix := 0; ix < stack_size; ix++ {
		ret, err := fsm.Pop()
		if ret != true {
			t.Error("FSM Pop error: failed to Pop a decoder from the FSM stack. /1")
		}
		if err != nil {
			t.Error("FSM Pop error: failed to Pop a decoder from the FSM stack. /2")
		}
		// compute expected size and compare
		expected_stack_size--
		cur_stack_size := fsm.StackSize()
		if expected_stack_size != cur_stack_size {
			t.Errorf("FSM Pop error: stack size [%d] different from that expected [%d].", cur_stack_size,
				expected_stack_size)
		}
	}
	// Now no more decoder should be present on the stack. Let's verify this.
	stack_size = fsm.StackSize()
	if stack_size != 0 {
		t.Error("FSM stack size error: stack is not empty.")
	}

	// Now calling Pop() should fail
	ret, err := fsm.Pop()
	if ret != false || err == nil {
		t.Error("FSM error: Pop didn't return an error on an empty stack.")
	}
}

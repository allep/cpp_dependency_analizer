// Implements unit tests for CppDecoder

package model

import (
	"fmt"
	"testing"
)

func TestCppDecoderIncludeDecodeLine(t *testing.T) {
	fmt.Println("Testing CppDecoderInclude > DecodeLine")
	var p CppDecoderInclude
	var tests = []struct {
		test_name string
		line_in string
		expected_error_out error
	}{
		{"Decode test 1", "#include <iostream>", nil},
		{"Decode test 2", "	#include <iostream>", nil},
		{"Decode test 3", "    #include <iostream>", nil},
		{"Decode test 4", "#include \"stdio.h\"", nil},
		{"Decode test 5", "	#include \"stdio.h\"", nil},
		{"Decode test 6", "    #include \"stdio.h\"", nil},
	}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			_, rerr := p.DecodeLine(tt.line_in)
			if rerr != tt.expected_error_out {
				t.Errorf("Decode test. Got an error different from the expected one [%v]", rerr)
			}
		})
	}
}

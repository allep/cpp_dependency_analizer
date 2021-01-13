// Implements unit tests for CppDecoder

package model

import (
	"fmt"
	"testing"
)

func TestCppDecoderGetTokensFromLine(t *testing.T) {
	fmt.Println("Testing GetTokenFromLine")
	var tests = []struct {
		test_name       string
		input_line      string
		requested_token int
		output_token    string
		is_error        bool
	}{
		{"Token test 1", "Sample sentence of tokens", 1, "sentence", false},
		{"Token test 2", "Sample sentence of tokens", -1, "", true},
		{"Token test 3", "  Sample sentence of tokens", 2, "of", false},
		{"Token test 4", "		Sample sentence of tokens", 3, "tokens", false},
	}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			str, err := GetTokenFromLine(tt.input_line, tt.requested_token)
			if tt.is_error && err == nil {
				t.Error("Expected an output error, instead got nil.")
			}
			if tt.is_error == false {
				if str != tt.output_token {
					t.Error("Expected token:[", tt.output_token, "], got: [", str, "]")
				}
			}
		})
	}
}

func TestCppDecoderIncludeDecodeLine(t *testing.T) {
	fmt.Println("Testing CppDecoderInclude > DecodeLine")
	var p CppDecoderInclude
	var tests = []struct {
		test_name          string
		line_in            string
		symbol_out         string
		expected_error_out bool
	}{
		{"Decode test 1", "#include <iostream1>", "iostream1", false},
		{"Decode test 2", "	#include <iostream2>", "iostream2", false},
		{"Decode test 3", "    #include <iostream3>", "iostream3", false},
		{"Decode test 4", "#include \"stdio1.h\"", "stdio1.h", false},
		{"Decode test 5", "	#include \"stdio2.h\"", "stdio2.h", false},
		{"Decode test 6", "    #include \"stdio3.h\"", "stdio3.h", false},
	}

	var expected_symbols []string
	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			_, rerr := p.DecodeLine(tt.line_in)
			// used inside next loop to compare expected vs extracted
			expected_symbols = append(expected_symbols, tt.symbol_out)
			if (tt.expected_error_out == true) && (rerr == nil) {
				t.Errorf("Decode test. Error condition different from the expected one.")
			}
		})
	}

	t.Log("Checking for symbol decoding mismatch")
	// now compare expected symbols with those actually extracted
	num_found := 0
	extacted_symbols := p.GetSymbols()
	for _, v1 := range expected_symbols {
		t.Log("Symbol decoding. Looking for:", v1)
		for _, v2 := range extacted_symbols {
			if v1 == v2 {
				t.Log("- Found:", v1)
				num_found++
			}
		}
	}
	// final check: have we found all expected symbols?
	if len(expected_symbols) != num_found {
		t.Error("Number of decoded symbol mismatch. Found:", num_found, ", expected:", len(expected_symbols))
	}
}

// Implements basic testing for CppTextParser

package model

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {
	// TODO: properly think about how to test this. Based on current interface this is not easily 
	// testable!

	fmt.Println("Testing ParseLine")
	var p CppTextParser
	p.ParseLine("Linea di prova")
}

func TestGetKeyFromLine(t *testing.T) {
	fmt.Println("Testing GetKeyFromLine")
	var p CppTextParser
	var tests = []struct {
		test_name string
		line_in string
		key_out string
	}{
		{"Include test 1", "#include <iostream>", "#include"},
		{"Typedef test 1", "typedef struct {", "typedef"},
		{"Typedef test 2", "	typedef struct foo_s {", "typedef"},
		{"Enum test 1", "enum foo_e {", "enum"},
		{"Enum test 2", "    enum foo_e {", "enum"},
		{"Class test 1", "class foo_c {", "class"},
		{"Class test 2", "    class foo_c {", "class"},
	}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			var k string
			k, _ = p.GetKeyFromLine(tt.line_in)
			if k != tt.key_out {
				t.Errorf("Key testing. Got %s, expected %s", k, tt.key_out)
			}
		})
	}
}

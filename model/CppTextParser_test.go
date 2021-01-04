// Implements basic testing for CppTextParser

package model

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {
	// TODO: in order to test this we need to test the whole functionality over a sample text
	// In other words, we should not try to test this against a single line, since if we choose to have
	// a generic interface we can't properly test it.

	fmt.Println("Testing CppTextParser > ParseLine")
	// TODO
}

func TestGetKeyFromLine(t *testing.T) {
	fmt.Println("Testing CppTextParser > GetKeyFromLine")
	var p CppTextParser
	var tests = []struct {
		test_name string
		line_in string
		key_out string
	}{
		{"Include test 1", "#include <iostream>", "#include"},
		{"Include test 2", "  #include <iostream>", "#include"},
		{"Include test 3", "	#include <iostream>", "#include"},
		{"Include test 4", " 	#include <iostream>", "#include"},
		{"Typedef test 1", "typedef struct {", "typedef"},
		{"Typedef test 2", "  typedef struct foo_s {", "typedef"},
		{"Typedef test 3", "	typedef struct foo_s {", "typedef"},
		{"Typedef test 4", " 	typedef struct foo_s {", "typedef"},
		{"Enum test 1", "enum foo_e {", "enum"},
		{"Enum test 2", "    enum foo_e {", "enum"},
		{"Enum test 3", "		enum foo_e {", "enum"},
		{"Enum test 4", " 		enum foo_e {", "enum"},
		{"Class test 1", "class foo_c {", "class"},
		{"Class test 2", "    class foo_c {", "class"},
		{"Class test 3", "		class foo_c {", "class"},
		{"Class test 4", " 		class foo_c {", "class"},
		{"Line comment test 1", "// this is a comment", "//"},
		{"Line comment test 2", "  // this is a comment", "//"},
		{"Line comment test 3", "	// this is a comment", "//"},
		{"Line comment test 4", "	 // this is a comment", "//"},
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

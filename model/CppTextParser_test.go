// Implements basic testing for CppTextParser

package model

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {
	fmt.Println("Testing CppTextParser > ParseLine")
	var p CppTextParser
	input_text := []string{
		"#include <iostream>",
		"namespace std;",
		"void main()",
		"{",
		"    int argn;",
		"    printf(\"Hello world!\");",
		"}",
		"",
		"class foo_c",
		"{",
		"    // comment",
		"    // inner class",
		"    class inner {",
		"        void inner_method();",
		"        uint16_t innerMember;",
		"    };",
		"    void method();",
		"    uint8_t aMember;",
		"};",
		"",
		"// #include <commented_include.h>",
		"#include \"../embedded_include.h\"",
		"// A double include",
		"#include <iostream>",
		"enum foo_e {",
		"    FOO_0,",
		"    BAR_1,",
		"};",
		"",
		"typedef struct foo_s {",
		"    uint8_t aField;",
		"} foo_t;",
		"// A useless include at the end of the file",
		"  #include \"../../useless_include.h\"",
	}

	p.Init()

	for _, v := range input_text {
		p.ParseLine(v)
	}

	// now let's get all symbols decoded from file
	includes := p.GetIncludes()

	// make sure we have decoded the correct number of includes
	num_includes := len(includes)
	if num_includes != 3 {
		t.Errorf("Number of includes [%d] does not matching with expected.", num_includes)
	}

	for _, i := range includes {
		fmt.Println("Found include:", i)
	}
}

func TestGetKeyFromLine(t *testing.T) {
	fmt.Println("Testing CppTextParser > GetKeyFromLine")
	var p CppTextParser
	var tests = []struct {
		test_name string
		line_in   string
		key_out   string
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

	p.Init()

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

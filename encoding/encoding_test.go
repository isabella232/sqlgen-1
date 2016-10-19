package encoding

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseArray(t *testing.T) {
	for _, tt := range []struct {
		input string
		delim string
		dims  []int
		elems [][]byte
	}{
		{`{}`, `,`, nil, [][]byte{}},
		{`{NULL}`, `,`, []int{1}, [][]byte{nil}},
		{`{a}`, `,`, []int{1}, [][]byte{{'a'}}},
		{`{a,b}`, `,`, []int{2}, [][]byte{{'a'}, {'b'}}},
		{`{{a,b}}`, `,`, []int{1, 2}, [][]byte{{'a'}, {'b'}}},
		{`{{a},{b}}`, `,`, []int{2, 1}, [][]byte{{'a'}, {'b'}}},
		{`{{{a,b},{c,d},{e,f}}}`, `,`, []int{1, 3, 2}, [][]byte{
			{'a'}, {'b'}, {'c'}, {'d'}, {'e'}, {'f'},
		}},
		{`{""}`, `,`, []int{1}, [][]byte{{}}},
		{`{","}`, `,`, []int{1}, [][]byte{{','}}},
		{`{",",","}`, `,`, []int{2}, [][]byte{{','}, {','}}},
		{`{{",",","}}`, `,`, []int{1, 2}, [][]byte{{','}, {','}}},
		{`{{","},{","}}`, `,`, []int{2, 1}, [][]byte{{','}, {','}}},
		{`{{{",",","},{",",","},{",",","}}}`, `,`, []int{1, 3, 2}, [][]byte{
			{','}, {','}, {','}, {','}, {','}, {','},
		}},
		{`{"\"}"}`, `,`, []int{1}, [][]byte{{'"', '}'}}},
		{`{"\"","\""}`, `,`, []int{2}, [][]byte{{'"'}, {'"'}}},
		{`{{"\"","\""}}`, `,`, []int{1, 2}, [][]byte{{'"'}, {'"'}}},
		{`{{"\""},{"\""}}`, `,`, []int{2, 1}, [][]byte{{'"'}, {'"'}}},
		{`{{{"\"","\""},{"\"","\""},{"\"","\""}}}`, `,`, []int{1, 3, 2}, [][]byte{
			{'"'}, {'"'}, {'"'}, {'"'}, {'"'}, {'"'},
		}},
		{`{axyzb}`, `xyz`, []int{2}, [][]byte{{'a'}, {'b'}}},
	} {
		dims, elems, err := parseArray([]byte(tt.input), []byte(tt.delim))

		if err != nil {
			t.Fatalf("Expected no error for %q, got %q", tt.input, err)
		}
		if !reflect.DeepEqual(dims, tt.dims) {
			t.Errorf("Expected %v dimensions for %q, got %v", tt.dims, tt.input, dims)
		}
		if !reflect.DeepEqual(elems, tt.elems) {
			t.Errorf("Expected %v elements for %q, got %v", tt.elems, tt.input, elems)
		}
	}
}

func TestParseArrayError(t *testing.T) {
	for _, tt := range []struct {
		input, err string
	}{
		{``, "expected '{' at offset 0"},
		{`x`, "expected '{' at offset 0"},
		{`}`, "expected '{' at offset 0"},
		{`{`, "expected '}' at offset 1"},
		{`{{}`, "expected '}' at offset 3"},
		{`{}}`, "unexpected '}' at offset 2"},
		{`{,}`, "unexpected ',' at offset 1"},
		{`{,x}`, "unexpected ',' at offset 1"},
		{`{x,}`, "unexpected '}' at offset 3"},
		{`{x,{`, "unexpected '{' at offset 3"},
		{`{x},`, "unexpected ',' at offset 3"},
		{`{x}}`, "unexpected '}' at offset 3"},
		{`{{x}`, "expected '}' at offset 4"},
		{`{""x}`, "unexpected 'x' at offset 3"},
		{`{{a},{b,c}}`, "multidimensional arrays must have elements with matching dimensions"},
	} {
		_, _, err := parseArray([]byte(tt.input), []byte{','})

		if err == nil {
			t.Fatalf("Expected error for %q, got none", tt.input)
		}
		if !strings.Contains(err.Error(), tt.err) {
			t.Errorf("Expected error to contain %q for %q, got %q", tt.err, tt.input, err)
		}
	}
}

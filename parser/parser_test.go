package parser

import (
	goparser "go/parser"
	"go/token"
	"io/ioutil"
	"testing"
)

var code []byte

var fields = map[string]struct{}{
	"ID":       struct{}{},
	"Created":  struct{}{},
	"Updated":  struct{}{},
	"Name":     struct{}{},
	"Email":    struct{}{},
	"Age":      struct{}{},
	"Hobbies":  struct{}{},
	"Colors":   struct{}{},
	"Address":  struct{}{},
	"emptyTag": struct{}{},
}

func validField(name string) bool {
	_, ok := fields[name]
	return ok
}

func TestInspectFields(t *testing.T) {
	code, err := ioutil.ReadFile("testcode")
	if err != nil {
		t.Fatal("testcode file missing")
	}

	fset := token.NewFileSet()
	f, err := goparser.ParseFile(fset, "test.go", code, 0)
	if err != nil {
		t.Fatal(err)
	}
	fields := inspectFields(f, "User")
	if fields == nil {
		t.Fatal("could not check existance")
	}
	for _, field := range fields {
		names := field.Names
		if names == nil {
			continue
		}
		for _, ident := range names {
			if !validField(ident.Name) {
				t.Fatalf("invalid field name, %q", ident.Name)
			}
		}
	}
}

func TestStructType(t *testing.T) {
	p, err := ParseFiles("testcode")
	if err != nil {
		t.Fatal("could not parse files:", err)
	}
	typ, ok := p.StructType("User")
	if !ok {
		t.Fatal("could not find struct")
	}
	for name := range typ.Fields {
		if !validField(name) {
			t.Fatal("unexpected field")
		}
	}
	f, _ := typ.Fields["Hobbies"]
	if f.Type != "[]Hobby" {
		t.Fatal("unexpected type for Hobbies:", f.Type)
	}
	f, _ = typ.Fields["Created"]
	if f.Type != "time.Time" {
		t.Fatal("unexpected type for Created:", f.Type)
	}
	f, _ = typ.Fields["Name"]
	if f.Type != "string" {
		t.Fatal("unexpected type for Created:", f.Type)
	}
}

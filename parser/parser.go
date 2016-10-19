package parser

import (
	"errors"
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	rx1 = regexp.MustCompile(`(.)([A-Z][a-z]+)`)
	rx2 = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

type Package struct {
	Name string

	files []*ast.File
}

// ParseDir parses the package in the given directory and returns it.
func ParseDir(directory string, skipSuffix string) (*Package, error) {
	dirFiles, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("cannto read directory %q: %s", directory, err)
	}

	var files []*ast.File
	fs := token.NewFileSet()
	for _, fd := range dirFiles {
		if fd.IsDir() {
			continue
		}
		name := fd.Name()
		if !strings.HasSuffix(name, ".go") ||
			(skipSuffix != "" && strings.HasSuffix(name, skipSuffix)) {
			continue
		}
		if directory != "." {
			name = filepath.Join(directory, name)
		}
		f, err := goparser.ParseFile(fs, name, nil, 0)
		if err != nil {
			log.Printf("ignoring unparsable file %q: %s", name, err)
			continue
		}
		files = append(files, f)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("%s: no buildable Go files", directory)
	}

	return &Package{
		Name:  files[0].Name.Name,
		files: files,
	}, nil
}

// ParseFiles parses the given files
func ParseFiles(filenames ...string) (*Package, error) {
	if filenames == nil {
		return nil, errors.New("must pass at least one file")
	}

	var files []*ast.File
	fs := token.NewFileSet()
	for _, name := range filenames {
		f, err := goparser.ParseFile(fs, name, nil, 0)
		if err != nil {
			return nil, fmt.Errorf("parsing file %v: %v", name, err)
		}
		files = append(files, f)
	}
	if len(files) == 0 {
		return nil, errors.New("no buildable Go files")
	}

	return &Package{
		Name:  files[0].Name.Name,
		files: files,
	}, nil
}

func (p *Package) Alias(name string) (*TypeAlias, bool) {
	if name == "" {
		return nil, false
	}
	t := TypeAlias{}
	var typeFound bool
	for _, f := range p.files {
		ast.Inspect(f, func(n ast.Node) bool {
			type_, ok := n.(*ast.TypeSpec)
			if !ok {
				if !typeFound {
					return true
				}
			} else {
				typeFound = true
				if type_.Name.Name != name {
					return false
				}
				i, ok := type_.Type.(*ast.Ident)
				if !ok {
					return false
				}
				t.Name = type_.Name.Name
				t.Type = i.Name
				return true
			}
			return true
		})
	}
	return &t, true
}

type TypeAlias struct {
	Name string
	Type string
}

// StructType returns a StructType if it exists in the package
func (p *Package) StructType(name string) (typ *StructType, ok bool) {
	if name == "" {
		return nil, false
	}
	for _, file := range p.files {
		possFields := inspectFields(file, name)
		if possFields != nil {
			return newStruct(name, possFields), true
		}
	}
	return nil, false
}

// StructType represents a go struct type
type StructType struct {
	TypeName string
	Fields   map[string]*Field
}

// inspectFields walks the ast tree to fetch the fields of the given structure
func inspectFields(f *ast.File, structName string) []*ast.Field {
	var fields []*ast.Field
	var typeFound bool
	ast.Inspect(f, func(n ast.Node) bool {
		type_, ok := n.(*ast.TypeSpec)
		if !ok {
			if !typeFound {
				return true
			}
		} else {
			typeFound = true
			if type_.Name.Name != structName {
				return false
			}
			return true
		}

		strt, ok := n.(*ast.StructType)
		if !ok {
			return true
		}
		fields = strt.Fields.List
		return true
	})
	return fields
}

func newStruct(name string, fields []*ast.Field) *StructType {
	typ := StructType{TypeName: name}
	typ.Fields = make(map[string]*Field)
	for _, field := range fields {
		if field.Names == nil {
			// Anonymous fields
			continue
		}
		for _, ident := range field.Names {
			// Ignore Function, methods and other type of fields
			if ident.Obj.Kind != ast.Var {
				continue
			}
			if ident.Name != "" {
				typ.Fields[ident.Name] = newField(ident.Name, field)
				continue
			}
		}
	}
	return &typ
}

func newField(name string, field *ast.Field) *Field {
	//TODO(sat): comments don't work yet, get it done
	var comments []string
	if field.Comment != nil {
		for _, c := range field.Comment.List {
			comments = append(comments, c.Text)
		}
	}
	var tags map[string][]string
	if field.Tag != nil {
		tags = decodeTag(field.Tag.Value)
	}

	return &Field{
		Name:    name,
		Type:    getType(field.Type),
		Tag:     tags,
		Comment: comments,
	}
}

func getType(f ast.Expr) string {
	// TODO: type for pointers
	switch v := f.(type) {
	case *ast.BasicLit:
		return v.Value
	case *ast.Ident:
		return v.Name
	case *ast.ArrayType:
		return "[]" + getType(v.Elt)
	case *ast.ChanType:
		return "chan " + getType(v.Value)
	case *ast.MapType:
		return "map[" + getType(v.Key) + "]" + getType(v.Value)
	case *ast.SelectorExpr:
		return getType(v.X) + "." + v.Sel.Name
	case *ast.StarExpr:
		return "*" + getType(v.X)
	default:
		return fmt.Sprintf("%T", v)
	}
}

func decodeTag(tag string) fieldTags {
	tag = strings.Trim(tag, "`")
	ret := make(fieldTags)
	for _, t := range strings.Split(tag, " ") {
		splits := strings.Split(t, ":")
		if len(splits) != 2 {
			// TODO(sat): should I panic on invalid tags?
			// Ignoring invalid tags
			continue
		}
		ret[splits[0]] = strings.Split(strings.Trim(splits[1], `"`), ",")
	}
	return ret
}

// Field represents a field in a struct
type Field struct {
	Name    string
	Type    string
	Tag     fieldTags
	Comment []string
}

type fieldTags map[string][]string

func (ft fieldTags) Has(tag, value string) bool {
	values, ok := ft[tag]
	if !ok {
		return false
	}
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func (f *Field) JSONName() string {
	jsont, ok := f.Tag["json"]
	if ok && len(jsont) > 0 {
		//TODO: what to do when it is `omit`?
		return jsont[0]
	}
	return f.Name
}

func (f *Field) IsArray() bool {
	return strings.HasPrefix(f.Type, "[]")
}

// ElemType return type name, skipping container part. For example, if the
// field type is []string, ElemType will return string.
func (f *Field) ElemType() string {
	tp := f.Type
	if f.IsArray() {
		tp = tp[2:]
	}
	for tp[0] == '*' {
		tp = tp[1:]
	}
	return tp
}

func (f *Field) IsPointer() bool {
	if f.IsArray() {
		return f.Type[2] == '*'
	}
	return f.Type[0] == '*'
}

// GoValueName return field name as should be used when assigned to variable
func (f *Field) GoValueName() string {
	name := f.Name
	return strings.ToLower(name[:1]) + name[1:]
}

type ConstValue struct {
	Name     string
	BaseType string
	Type     string
	Value    interface{}
	FitInt32 bool
}

// JSONValue return const value as should be used in JSON format
func (cv *ConstValue) JSONValue() string {
	if strings.HasPrefix(cv.Name, cv.Type) {
		return strings.TrimLeft(cv.Name, cv.Type)
	}
	return cv.Name
}

// ConstValue return all const declarations
func (p *Package) ConstValues() map[string][]*ConstValue {
	consts := make(map[string][]*ConstValue, 0)
	types := make(map[string]string)

	for _, f := range p.files {
		ast.Inspect(f, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.Ident:
				if n.Obj == nil {
					return true
				}
				if n.Obj.Kind == ast.Typ {
					if tspec, ok := n.Obj.Decl.(*ast.TypeSpec); ok {
						types[n.Name] = fmt.Sprint(tspec.Type)
					}
				}
				if n.Obj.Kind != ast.Con { // ast.Con == constant
					return true
				}
				spec, ok := n.Obj.Decl.(*ast.ValueSpec)
				if !ok {
					return true
				}
				specType := fmt.Sprint(spec.Type)

				consts[specType] = append(consts[specType], &ConstValue{
					BaseType: types[specType],
					Type:     specType,
					Name:     n.Name,
					Value:    spec.Values[0].(*ast.BasicLit).Value,
				})
			}
			return true
		})
	}

	return consts
}

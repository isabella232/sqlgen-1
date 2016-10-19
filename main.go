package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/optiopay/generate-composite-types/parser"
)

var aliastypeTmpl, sqltypeTmpl, createSQLTmpl *template.Template

func init() {
	dataAlias, err := Asset("templates/aliastype.go.tmpl")
	if err != nil {
		log.Fatal("template was not embedded, run go-bindata")
	}
	aliastypeTmpl = template.Must(template.New("aliastype").Parse(string(dataAlias)))

	dataSql, err := Asset("templates/sqltype.go.tmpl")
	if err != nil {
		log.Fatal("template was not embedded, run go-bindata")
	}
	sqltypeTmpl = template.Must(template.New("sqltype").Parse(string(dataSql)))

	dataCreate, err := Asset("templates/create-type.sql.tmpl")
	if err != nil {
		log.Fatal("template was not embedded, run go-bindata")
	}
	createSQLTmpl = template.Must(template.New("createsql").Parse(string(dataCreate)))
}

func main() {
	log.SetFlags(log.Lshortfile)
	path := flag.String("path", ".", "path for the go source files")
	structs := flag.String("structs", "", "comma separated structs that needs generation")
	alias := flag.String("alias", "", "comma separated type aliases that needs generation")
	array := flag.Bool("array", false, "generate Array types")
	sql := flag.Bool("sql", false, "write sql initalizing functions")
	flag.Parse()

	g, err := NewGenerator(*path)
	fatalOnErr(err)

	if *structs != "" {
		for _, ty := range strings.Split(*structs, ",") {
			err = g.GenerateStruct(*path, ty, *sql, *array)
			fatalOnErr(err)
		}
	}

	if *alias != "" {
		for _, ty := range strings.Split(*alias, ",") {
			err = g.GenerateAlias(*path, ty)
			fatalOnErr(err)
		}
	}
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Generator struct {
	pkg *parser.Package
}

func NewGenerator(path string) (*Generator, error) {
	pkg, err := parser.ParseDir(path, "")
	if err != nil {
		return nil, fmt.Errorf("cannot parse %q: %s", path, err)
	}
	g := &Generator{
		pkg: pkg,
	}
	return g, nil
}

func (g *Generator) GenerateStruct(path, struct_ string, sql, array bool) error {
	st, ok := g.pkg.StructType(struct_)
	if !ok {
		return fmt.Errorf("%s type not found", struct_)
	}
	sqlTy, err := NewType(g.pkg, st)
	if err != nil {
		return err
	}
	filename := filepath.Join(path, strings.ToLower(struct_)+"_sql.go")
	var buf bytes.Buffer
	err = g.WriteType(sqlTy, &buf, struct_, array)
	if err != nil {
		return err
	}
	err = cleanupAndWrite(&buf, filename)
	if err != nil {
		return err
	}

	if !sql {
		return nil
	}
	f, err := os.OpenFile(filepath.Join(path, strings.ToLower(struct_)+".sql"),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return g.WriteSQL(sqlTy, f, struct_)
}

func (g *Generator) WriteSQL(type_ *Type, w io.Writer, ty string) error {
	return createSQLTmpl.Execute(w, map[string]interface{}{
		"SQLTypeName": strings.ToLower(type_.Name) + "_type",
		"Fields":      type_.Fields,
	})
}

func (g *Generator) WriteType(type_ *Type, w io.Writer, ty string, array bool) error {
	return sqltypeTmpl.Execute(w, map[string]interface{}{
		"TypeName":    type_.Name,
		"SQLTypeName": strings.ToLower(type_.Name) + "_type",
		"PackageName": g.pkg.Name,
		"Fields":      type_.Fields,
		"Array":       array,
	})
}

func (g *Generator) GenerateAlias(path, alias string) error {
	filename := filepath.Join(path, strings.ToLower(alias)+"_sql.go")
	var buf bytes.Buffer
	err := g.WriteAlias(&buf, alias)
	if err != nil {
		return err
	}
	return cleanupAndWrite(&buf, filename)
}

func (g *Generator) WriteAlias(w io.Writer, name string) error {
	alias, ok := g.pkg.Alias(name)
	if !ok {
		return fmt.Errorf("could not find type")
	}
	scanner, err := generateAliasScanner(alias)
	if err != nil {
		return err
	}
	scanBytes, err := generateAliasScanBytes(alias)
	if err != nil {
		return err
	}
	valuer, err := generateAliasValuer(alias)
	if err != nil {
		return err
	}
	return aliastypeTmpl.Execute(w, map[string]interface{}{
		"AliasName":   alias.Name,
		"PackageName": g.pkg.Name,
		"Scanner":     scanner,
		"ScanBytes":   scanBytes,
		"Valuer":      valuer,
	})
}

func cleanupAndWrite(b *bytes.Buffer, filename string) error {
	// cleanup and add imports and format code
	cleaned, err := imports.Process(filename, b.Bytes(), nil)
	if err != nil {
		log.Print("could not format and add imports, please use goimports and gofmt on the generated code")
		cleaned = b.Bytes()
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(cleaned)
	return err
}

func generateAliasScanner(a *parser.TypeAlias) (string, error) {
	resp := ""
	switch a.Type {
	case "int8", "int16", "int32", "int", "int64", "uint8", "uint16", "uint32", "uint", "uint64":
		resp = fmt.Sprintf(`case int64:
	        *x = %s(src)
	        return nil `, a.Name)
	case "float32", "float64":
		resp = fmt.Sprintf(`case float64:
	        *x = %s(src)
	        return nil `, a.Name)
	case "bool":
		resp = fmt.Sprintf(`case bool:
	        *x = %s(src)
	        return nil `, a.Name)
	default:
		return "", fmt.Errorf("unsupported type for %s: %s", a.Name, a.Type)
	}
	return resp, nil
}

func generateAliasScanBytes(a *parser.TypeAlias) (string, error) {
	resp := ""
	switch a.Type {
	case "string":
		resp = fmt.Sprintf(`*x = string(src)`)
	case "int8":
		resp = fmt.Sprintf(`i, err := strconv.ParseInt(string(src), 10, 8)
			if err != nil {
			return err
		}`)
	case "int16":
		resp = fmt.Sprintf(`i, err := strconv.ParseInt(string(src), 10, 16)
			if err != nil {
			return err
		}`)
	case "int32":
		resp = fmt.Sprintf(`i, err := strconv.ParseInt(string(src), 10, 32)
			if err != nil {
			return err
		}`)
	case "int":
		resp = fmt.Sprintf(`i, err := strconv.ParseInt(string(src), 10, 32)
			if err != nil {
			return err
		}`)
	case "int64":
		resp = fmt.Sprintf(`i, err := strconv.ParseInt(string(src), 10, 64)
			if err != nil {
			return err
		}`)
	case "uint8":
		resp = fmt.Sprintf(`i, err := strconv.ParseUint(string(src), 10, 8)
			if err != nil {
			return err
		}`)
	case "uint16":
		resp = fmt.Sprintf(`i, err := strconv.ParseUint(string(src), 10, 16)
			if err != nil {
			return err
		}`)
	case "uint32":
		resp = fmt.Sprintf(`i, err := strconv.ParseUint(string(src), 10, 32)
			if err != nil {
			return err
		}`)
	case "uint":
		resp = fmt.Sprintf(`i, err := strconv.ParseUint(string(src), 10, 32)
			if err != nil {
			return err
		}`)
	case "uint64":
		resp = fmt.Sprintf(`i, err := strconv.ParseUint(string(src), 10, 64)
			if err != nil {
			return err
		}`)
	case "float32":
		resp = fmt.Sprintf(`i, err := strconv.ParseFloat(string(src), 32)
			if err != nil {
			return err
		}`)
	case "float64":
		resp = fmt.Sprintf(`i, err := strconv.ParseFloat(string(src), 64)
			if err != nil {
			return err
		}`)
	case "bool":
		resp = fmt.Sprintf(`
		var i bool
		switch string(src) {
		case "t":
			i = true
		case "f":
			i = false
		default:
			return fmt.Errorf("unexpected value for bool")
		} `)
	default:
		return "", fmt.Errorf("unsupported type for %s: %s", a.Name, a.Type)
	}
	return resp, nil
}

func generateAliasValuer(a *parser.TypeAlias) (string, error) {
	resp := ""
	switch a.Type {
	case "int", "int8", "int16", "int32", "int64":
		resp = fmt.Sprintf(`b = strconv.AppendInt(b, int64(x), 10)`)
	case "uint", "uint8", "uint16", "uint32", "uint64":
		resp = fmt.Sprintf(`b = strconv.AppendUint(b, uint64(x), 10)`)
	case "float32":
		resp = fmt.Sprintf(`b = strconv.AppendFloat(b, float64(x), 'f', -1, 32)`)
	case "float64":
		resp = fmt.Sprintf(`b = strconv.AppendFloat(b, float64(x), 'f', -1, 64)`)
	case "string":
		resp = fmt.Sprintf(`b = encoding.AppendArrayQuotedBytes(b, []byte(x))`)
	case "bool":
		resp = fmt.Sprintf(`b = strconv.AppendBool(b, x)`)
	default:
		return "", fmt.Errorf("unsupported alias for type %s: %s", a.Name, a.Type)
	}
	return resp, nil
}

type Type struct {
	Name   string
	Fields FieldArray
}

type Field struct {
	Name    string
	Type    string
	SQLName string
	SQLType string
}

type FieldArray []Field

func (f FieldArray) Len() int           { return len(f) }
func (f FieldArray) Less(i, j int) bool { return f[i].Name < f[j].Name }
func (f FieldArray) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func NewType(pkg *parser.Package, st *parser.StructType) (*Type, error) {
	sql := &Type{
		Name: st.TypeName,
	}
	for _, f := range st.Fields {
		goType := f.Type
		// Remove * from the name
		if f.IsPointer() {
			goType = goType[1:]
		}
		field := Field{
			Name:    f.Name,
			Type:    f.Type,
			SQLName: strings.ToLower(f.Name),
			SQLType: determineSQLType(pkg, goType),
		}
		sql.Fields = append(sql.Fields, field)
	}
	sort.Sort(sql.Fields)
	return sql, nil
}

func determineSQLType(pkg *parser.Package, goType string) string {
	var ty string
	switch goType {
	case "int8", "uint8", "int16", "uint16":
		ty = "smallint"
	case "int", "int32", "uint32", "uint":
		ty = "integer"
	case "int64", "uint64":
		ty = "bigint"
	case "float32":
		ty = "real"
	case "float64":
		ty = "double precision"
	case "string":
		ty = "text"
	case "bool":
		ty = "boolean"
	case "time.Time":
		ty = "timestamp with time zone"
	default:
		alias, ok := pkg.Alias(goType)
		if ok && alias.Name != "" {
			ty = determineSQLType(pkg, alias.Type)
		} else {
			// if the type belongs to another package
			chunks := strings.Split(goType, ".")
			typ := chunks[len(chunks)-1]
			ty = strings.ToLower(typ) + "_type"
		}
	}
	return ty
}

func (t *Type) FieldsAsStrings() string {
	fields := make([]string, 0, len(t.Fields))
	for _, f := range t.Fields {
		fields = append(fields, strings.ToLower(f.Name)+" "+f.SQLType)
	}
	return strings.Join(fields, ",")
}

func (f *Field) Valuer(i int) string {
	var resp string
	switch f.Type {
	case "int", "int8", "int16", "int32", "int64":
		resp = fmt.Sprintf(`b = strconv.AppendInt(b, int64(x.%s), 10)`, f.Name)
	case "uint", "uint8", "uint16", "uint32", "uint64":
		resp = fmt.Sprintf(`b = strconv.AppendUint(b, uint64(x.%s), 10)`, f.Name)
	case "float32":
		resp = fmt.Sprintf(`b = strconv.AppendFloat(b, float64(x.%s), 'f', -1, 32)`, f.Name)
	case "float64":
		resp = fmt.Sprintf(`b = strconv.AppendFloat(b, float64(x.%s), 'f', -1, 64)`, f.Name)
	case "string":
		resp = fmt.Sprintf(`b = encoding.AppendArrayQuotedBytes(b, []byte(x.%s))`, f.Name)
	case "bool":
		resp = fmt.Sprintf(`b = strconv.AppendBool(b, x.%s)`, f.Name)
	case "time.Time":
		resp = fmt.Sprintf(`b = append(b, pq.FormatTimestamp(x.%s)...)`, f.Name)
	case "*time.Time":
		resp = fmt.Sprintf(`b = append(b, pq.FormatTimestamp(*(x.%s))...)`, f.Name)
	default:
		resp = fmt.Sprintf(`resp%d, err := x.%s.Value()
	if err != nil {
		return nil, err
	}
	raw%d, ok := resp%d.([]byte)
	if !ok {
		return nil, fmt.Errorf("unexpected error")
	}
	b = encoding.AppendArrayQuotedBytes(b, raw%d)`, i, f.Name, i, i, i)
	}
	return resp
}

func (f *Field) Scanner(i int) string {
	var resp string
	switch f.Type {
	case "string":
		resp = fmt.Sprintf(`x.%s = string(elems[%d])`, f.Name, i)
	case "int16":
		resp = fmt.Sprintf(`x%d, err := strconv.ParseInt(string(elems[%d]), 10, 16)
			if err != nil {
			return err
		}
		x.%s = int16(x%d)`, i, i, f.Name, i)
	case "int32":
		resp = fmt.Sprintf(`x%d, err := strconv.ParseInt(string(elems[%d]), 10, 32)
			if err != nil {
			return err
		}
		x.%s = int32(x%d)`, i, i, f.Name, i)
	case "int":
		resp = fmt.Sprintf(`x%d, err := strconv.ParseInt(string(elems[%d]), 10, 32)
			if err != nil {
			return err
		}
		x.%s = int(x%d)`, i, i, f.Name, i)
	case "int64":
		resp = fmt.Sprintf(`x.%s, err = strconv.ParseInt(string(elems[%d]), 10, 64)
			if err != nil {
			return err
		}`, f.Name, i)
	case "uint16":
		resp = fmt.Sprintf(`x%d, err := strconv.ParseUint(string(elems[%d]), 10, 16)
			if err != nil {
			return err
		}
		x.%s = uint16(x%d)`, i, i, f.Name, i)
	case "uint32":
		resp = fmt.Sprintf(`x%d, err := strconv.ParseUint(string(elems[%d]), 10, 32)
			if err != nil {
			return err
		}
		x.%s = uint32(x%d)`, i, i, f.Name, i)
	case "uint":
		resp = fmt.Sprintf(`x%d, err := strconv.ParseUint(string(elems[%d]), 10, 32)
			if err != nil {
			return err
		}
		x.%s = uint(x%d)`, i, i, f.Name, i)
	case "uint64":
		resp = fmt.Sprintf(`x.%s, err = strconv.ParseUint(string(elems[%d]), 10, 64)
			if err != nil {
			return err
		}`, f.Name, i)

	case "float32":
		resp = fmt.Sprintf(`x%d, err := strconv.ParseFloat(string(elems[%d]), 32)
			if err != nil {
			return err
		}
		x.%s = float32(x%d)`, i, i, f.Name, i)
	case "float64":
		resp = fmt.Sprintf(`x.%s, err = strconv.ParseFloat(string(elems[%d]), 64)
			if err != nil {
			return err
		}`, f.Name, i)
	case "bool":
		resp = fmt.Sprintf(`switch string(elems[%d]) {
		case "t":
			x.%s = true
		case "f":
			x.%s = false
		default:
			return fmt.Errorf("unexpected value for bool")
		} `, i, f.Name, f.Name)
	case "time.Time":
		resp = fmt.Sprintf(`t%d, err := pq.ParseTimestamp(time.UTC, string(elems[%d]))
			if err != nil {
				return err
			}
			x.%s = t%d`, i, i, f.Name, i)
	case "*time.Time":
		resp = fmt.Sprintf(`t%d, err := pq.ParseTimestamp(time.UTC, string(elems[%d]))
			if err != nil {
				return err
			}
			x.%s = &t%d`, i, i, f.Name, i)
	default:
		if strings.HasPrefix(f.Type, "*") {
			resp = fmt.Sprintf(`var y%d = new(%s)`, i, strings.Trim(f.Type, "*"))
		} else {
			resp = fmt.Sprintf(`var y%d %s`, i, strings.Trim(f.Type, "*"))
		}
		resp += fmt.Sprintf(`
			err = y%d.Scan(elems[%d])
			if err != nil {
				return err
			}
			x.%s = y%d `, i, i, f.Name, i)
	}
	return resp
}

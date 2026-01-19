package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type introspectionResult struct {
	Data struct {
		Schema schema `json:"__schema"`
	} `json:"data"`
}

type schema struct {
	Types []introspectionType `json:"types"`
}

type introspectionType struct {
	Kind        string                  `json:"kind"`
	Name        string                  `json:"name"`
	Description *string                 `json:"description"`
	Fields      []introspectionField    `json:"fields"`
	InputFields []introspectionInputVal `json:"inputFields"`
	EnumValues  []introspectionEnumVal  `json:"enumValues"`
}

type introspectionField struct {
	Name string               `json:"name"`
	Type introspectionTypeRef `json:"type"`
}

type introspectionInputVal struct {
	Name string               `json:"name"`
	Type introspectionTypeRef `json:"type"`
}

type introspectionEnumVal struct {
	Name string `json:"name"`
}

type introspectionTypeRef struct {
	Kind   string                `json:"kind"`
	Name   *string               `json:"name"`
	OfType *introspectionTypeRef `json:"ofType"`
}

var objectTypeNames map[string]bool
var inputTypeNames map[string]bool
var enumTypeNames map[string]bool
var scalarTypeNames map[string]bool

func main() {
	schemaData, err := os.ReadFile("sui_mainnet_schema.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read schema: %v\n", err)
		os.Exit(1)
	}
	var result introspectionResult
	err = json.Unmarshal(schemaData, &result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal schema: %v\n", err)
		os.Exit(1)
	}
	code, err := generateTypes(result.Data.Schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate types: %v\n", err)
		os.Exit(1)
	}
	err = os.WriteFile("sui_types.go", code, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write types: %v\n", err)
		os.Exit(1)
	}
}

func generateTypes(s schema) ([]byte, error) {
	objectTypeNames = map[string]bool{}
	inputTypeNames = map[string]bool{}
	enumTypeNames = map[string]bool{}
	scalarTypeNames = map[string]bool{}
	var objects []introspectionType
	var inputs []introspectionType
	var enums []introspectionType
	for _, t := range s.Types {
		if strings.HasPrefix(t.Name, "__") {
			continue
		}
		switch t.Kind {
		case "OBJECT":
			objects = append(objects, t)
			objectTypeNames[t.Name] = true
		case "INPUT_OBJECT":
			inputs = append(inputs, t)
			inputTypeNames[t.Name] = true
		case "ENUM":
			enums = append(enums, t)
			enumTypeNames[t.Name] = true
		case "SCALAR":
			scalarTypeNames[t.Name] = true
		}
	}
	sort.Slice(objects, func(i, j int) bool { return objects[i].Name < objects[j].Name })
	sort.Slice(inputs, func(i, j int) bool { return inputs[i].Name < inputs[j].Name })
	sort.Slice(enums, func(i, j int) bool { return enums[i].Name < enums[j].Name })
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "package sui")
	fmt.Fprintln(&buf)
	for _, t := range enums {
		if len(t.EnumValues) == 0 {
			continue
		}
		fmt.Fprintf(&buf, "type %s string\n\n", t.Name)
		for _, v := range t.EnumValues {
			fmt.Fprintf(&buf, "const %s_%s %s = \"%s\"\n", t.Name, toConstSuffix(v.Name), t.Name, v.Name)
		}
		fmt.Fprintln(&buf)
	}
	for _, t := range objects {
		writeStruct(&buf, t.Name, t.Fields)
	}
	for _, t := range inputs {
		writeInputStruct(&buf, t.Name, t.InputFields)
	}
	return buf.Bytes(), nil
}

func writeStruct(buf *bytes.Buffer, name string, fields []introspectionField) {
	fmt.Fprintf(buf, "type %s struct {\n", name)
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	for _, f := range fields {
		goName := toExported(f.Name)
		goType := goTypeFromRef(&f.Type)
		fmt.Fprintf(buf, "\t%s %s `json:\"%s\"`\n", goName, goType, f.Name)
	}
	fmt.Fprintln(buf, "}\n")
}

func writeInputStruct(buf *bytes.Buffer, name string, fields []introspectionInputVal) {
	fmt.Fprintf(buf, "type %s struct {\n", name)
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	for _, f := range fields {
		goName := toExported(f.Name)
		goType := goTypeFromRef(&f.Type)
		fmt.Fprintf(buf, "\t%s %s `json:\"%s\"`\n", goName, goType, f.Name)
	}
	fmt.Fprintln(buf, "}\n")
}

func toExported(name string) string {
	if name == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(name)
	return string(unicode.ToUpper(r)) + name[size:]
}

func toConstSuffix(name string) string {
	var b strings.Builder
	for i, r := range name {
		if r == '-' || r == ' ' {
			b.WriteRune('_')
			continue
		}
		if unicode.IsLower(r) && i == 0 {
			b.WriteRune(unicode.ToUpper(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func goTypeFromRef(ref *introspectionTypeRef) string {
	named, listDepth := unwrapType(ref)
	base := goBaseType(named)
	if objectTypeNames[named] || inputTypeNames[named] {
		base = "*" + base
	}
	for i := 0; i < listDepth; i++ {
		base = "[]" + base
	}
	return base
}

func unwrapType(ref *introspectionTypeRef) (string, int) {
	depth := 0
	r := ref
	for r != nil && r.Kind == "NON_NULL" {
		r = r.OfType
	}
	for r != nil && r.Kind == "LIST" {
		depth++
		r = r.OfType
		for r != nil && r.Kind == "NON_NULL" {
			r = r.OfType
		}
	}
	for r != nil && r.OfType != nil {
		r = r.OfType
	}
	if r != nil && r.Name != nil {
		return *r.Name, depth
	}
	return "interface{}", depth
}

func goBaseType(name string) string {
	if name == "" {
		return "interface{}"
	}
	switch name {
	case "String", "ID":
		return "string"
	case "Boolean":
		return "bool"
	case "Int":
		return "int"
	case "Float":
		return "float64"
	}
	if scalarTypeNames[name] {
		return "string"
	}
	if enumTypeNames[name] {
		return name
	}
	if objectTypeNames[name] || inputTypeNames[name] {
		return name
	}
	return "interface{}"
}

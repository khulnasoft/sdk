// Copyright 2016-2018, KhulnaSoft, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:lll
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/khulnasoft/sdk/go/common/slice"
)

type builtin struct {
	Name           string
	Type           string
	inputType      string
	implements     []string
	Implements     []*builtin
	elementType    string
	item           *builtin
	ItemType       string
	Example        string
	elemExample    string
	GenerateConfig bool
	DefaultConfig  string
	RegisterInput  bool
	defaultValue   string

	// Used with the "array-contravariance" strategy to attempt to convert []interface{} slices to
	// []T.
	InnerElementType string

	Strategy string
}

const AssetOrArchiveType = "AssetOrArchive"

func (b builtin) DefineInputType() bool {
	return b.inputType == "" && b.Type != AssetOrArchiveType
}

func (b builtin) DefinePtrType() bool {
	return strings.HasSuffix(b.Name, "Ptr")
}

func (b builtin) PtrType() string {
	return b.inputType[1:]
}

func (b builtin) DefineInputMethods() bool {
	return b.Type != AssetOrArchiveType
}

func (b builtin) ImplementsPtrType() bool {
	switch b.Type {
	case "bool", "float64", "ID", "int", "string", "URN":
		return true
	}
	return false
}

func (b builtin) DefaultValue() string {
	if b.defaultValue != "" {
		return b.defaultValue
	}
	return b.Name + "{}"
}

func (b builtin) DefineElem() bool {
	return b.DefinePtrType()
}

func (b builtin) ElemReturnType() string {
	return strings.TrimSuffix(b.Name, "Ptr")
}

func (b builtin) ElemElementType() string {
	return strings.TrimPrefix(b.Type, "*")
}

func (b builtin) DefineIndex() bool {
	return strings.HasSuffix(b.Name, "Array")
}

func (b builtin) IndexReturnType() string {
	return strings.TrimSuffix(b.Name, "Array")
}

func (b builtin) IndexElementType() string {
	return strings.TrimPrefix(b.elementType, "[]")
}

func (b builtin) DefineMapIndex() bool {
	return strings.HasSuffix(b.Name, "Map")
}

func (b builtin) MapIndexElementType() string {
	return strings.TrimPrefix(b.elementType, "map[string]")
}

func (b builtin) MapIndexReturnType() string {
	return strings.TrimSuffix(b.Name, "Map")
}

func (b builtin) IndexConversion() string {
	toType := b.item.Name
	switch {
	case b.item.inputType != "":
		// No conversion necessary for types which are their own input type.
		return ""
	case toType == "Input":
		// There is no direct conversion into `Input`, so special case to use `ToOutput`
		return "ToOutput"
	case b.item.item == nil:
		// The item type is a primitive - so use a direct conversion into the defined type (e.g. `String(...)`)
		return toType
	default:
		// The item type is itself a container - so use one of the other code-generated `To<Type>` methods.
		return "To" + toType
	}
}

func (b builtin) ElementType() string {
	if b.elementType != "" {
		return b.elementType
	}
	return b.Type
}

func (b builtin) InputType() string {
	if b.inputType != "" {
		return b.inputType
	}
	return b.Name
}

func (b builtin) DefineToFunction() bool {
	if b.item != nil {
		return b.item.DefineToFunction()
	}
	return b.DefineInputMethods()
}

func (b builtin) ItemExample() string {
	return b.item.Example
}

func (b builtin) ElemExample() string {
	if strings.HasSuffix(b.Name, "Map") {
		return fmt.Sprintf("{\"baz\": %s}", b.item.ElemExample())
	}
	if strings.HasSuffix(b.Name, "Array") {
		return fmt.Sprintf("{%s}", b.item.ElemExample())
	}
	if b.elemExample != "" {
		return b.elemExample
	}
	return b.Example
}

var builtins = makeBuiltins([]*builtin{
	{Name: "Archive", Type: "Archive", inputType: "*archive", implements: []string{AssetOrArchiveType}, Example: "NewFileArchive(\"foo.zip\")"},
	{Name: "Asset", Type: "Asset", inputType: "*asset", implements: []string{AssetOrArchiveType}, Example: "NewFileAsset(\"foo.txt\")"},
	{Name: AssetOrArchiveType, Type: AssetOrArchiveType, Example: "NewFileArchive(\"foo.zip\")"},
	{Name: "Bool", Type: "bool", Example: "Bool(true)", GenerateConfig: true, DefaultConfig: "false", elemExample: "true", RegisterInput: true, defaultValue: "Bool(false)"},
	{Name: "Float64", Type: "float64", Example: "Float64(999.9)", GenerateConfig: true, DefaultConfig: "0", elemExample: "999.9", RegisterInput: true, defaultValue: "Float64(0)"},
	{Name: "ID", Type: "ID", inputType: "ID", implements: []string{"String"}, Example: "ID(\"foo\")", RegisterInput: true, defaultValue: "ID(\"\")"},
	{Name: "Input", Type: "interface{}", Example: "String(\"any\")"},
	{Name: "Int", Type: "int", Example: "Int(42)", GenerateConfig: true, DefaultConfig: "0", elemExample: "42", RegisterInput: true, defaultValue: "Int(0)"},
	{Name: "String", Type: "string", Example: "String(\"foo\")", elemExample: "\"foo\"", RegisterInput: true, defaultValue: "String(\"\")"},
	{Name: "URN", Type: "URN", inputType: "URN", implements: []string{"String"}, Example: "URN(\"foo\")", RegisterInput: true, defaultValue: "URN(\"\")"},
})

func unexported(s string) string {
	runes := []rune(s)

	allCaps := true
	for _, r := range runes {
		if !unicode.IsUpper(r) {
			allCaps = false
			break
		}
	}

	if allCaps {
		return strings.ToLower(s)
	}
	return string(append([]rune{unicode.ToLower(runes[0])}, runes[1:]...))
}

var funcs = template.FuncMap{
	"Unexported": unexported,
}

func makeBuiltins(primitives []*builtin) []*builtin {
	// Augment primitives with array and map types.
	builtins := slice.Prealloc[*builtin](len(primitives))
	for _, p := range primitives {
		p.Strategy = "primitive"
		name := ""
		if p.Name != "Input" {
			builtins = append(builtins, p)
			name = p.Name
		}
		switch name {
		case "Archive", "Asset", AssetOrArchiveType, "":
			// do nothing
		default:
			builtins = append(builtins, &builtin{
				Name:          name + "Ptr",
				Type:          "*" + p.Type,
				inputType:     "*" + unexported(p.Type) + "Ptr",
				Example:       fmt.Sprintf("%sPtr(%s(%s))", name, p.Type, p.Example),
				RegisterInput: p.RegisterInput,
				defaultValue:  p.defaultValue,
			})
		}
		arrType := &builtin{
			Name:          name + "Array",
			Type:          "[]" + name + "Input",
			ItemType:      name + "Input",
			elementType:   "[]" + p.Type,
			item:          p,
			Example:       fmt.Sprintf("%sArray{%s}", name, p.Example),
			RegisterInput: true,

			InnerElementType: p.Type,
		}
		builtins = append(builtins, arrType)
		mapType := &builtin{
			Name:          name + "Map",
			Type:          "map[string]" + name + "Input",
			ItemType:      name + "Input",
			elementType:   "map[string]" + p.Type,
			item:          p,
			Example:       fmt.Sprintf("%sMap{\"baz\": %s}", name, p.Example),
			RegisterInput: true,
		}
		builtins = append(builtins, mapType)
		builtins = append(builtins, &builtin{
			Name:          name + "ArrayMap",
			Type:          "map[string]" + name + "ArrayInput",
			ItemType:      name + "ArrayInput",
			elementType:   "map[string][]" + p.Type,
			item:          arrType,
			Example:       fmt.Sprintf("%sArrayMap{\"baz\": %sArray{%s}}", name, name, p.Example),
			RegisterInput: true,
		})
		builtins = append(builtins, &builtin{
			Name:          name + "MapArray",
			Type:          "[]" + name + "MapInput",
			ItemType:      name + "MapInput",
			elementType:   "[]map[string]" + p.Type,
			item:          mapType,
			Example:       fmt.Sprintf("%sMapArray{%sMap{\"baz\": %s}}", name, name, p.Example),
			RegisterInput: true,
		})
		mapMapType := &builtin{
			Name:          name + "MapMap",
			Type:          "map[string]" + name + "MapInput",
			ItemType:      name + "MapInput",
			elementType:   "map[string]map[string]" + p.Type,
			item:          mapType,
			Example:       fmt.Sprintf("%sMapMap{\"baz\": %sMap{\"baz\": %s}}", name, name, p.Example),
			RegisterInput: true,
		}
		builtins = append(builtins, mapMapType)
		arrayArrayType := &builtin{
			Name:          name + "ArrayArray",
			Type:          "[]" + name + "ArrayInput",
			ItemType:      name + "ArrayInput",
			elementType:   "[][]" + p.Type,
			item:          arrType,
			Example:       fmt.Sprintf("%sArrayArray{%sArray{%s}}", name, name, p.Example),
			RegisterInput: true,
		}
		builtins = append(builtins, arrayArrayType)

		// TODO - consider expanding to all primitives?
		if name == "" {
			builtins = append(builtins, &builtin{
				Name:          "ArrayArrayMap",
				Type:          "map[string]ArrayArrayInput",
				ItemType:      "ArrayArrayInput",
				elementType:   "map[string][][]" + p.Type,
				item:          arrayArrayType,
				Example:       fmt.Sprintf("%sArrayArrayMap{\"baz\": %sArrayArray{Array{%s}}}", name, name, p.Example),
				RegisterInput: true,
			})
		}

		// Unblock https://github.com/khulnasoft/khulnasoft/issues/17415
		if name == "String" {
			builtins = append(builtins, &builtin{
				Name:          name + "MapMapMap",
				Type:          "map[string]" + name + "MapMapInput",
				ItemType:      name + "MapMapInput",
				elementType:   "map[string]map[string]map[string]" + p.Type,
				item:          mapMapType,
				Example:       fmt.Sprintf("%sMapMapMap{\"baz\": %sMapMap{\"baz\": %sMap{\"baz\": %s}}}", name, name, name, p.Example),
				RegisterInput: true,
			})
		}
	}

	nameToBuiltin := map[string]*builtin{}
	for _, b := range builtins {
		nameToBuiltin[b.Name] = b
	}

	for _, b := range builtins {
		for _, i := range b.implements {
			b.Implements = append(b.Implements, nameToBuiltin[i])
		}
	}

	return builtins
}

func main() {
	templates, err := template.New("templates").Funcs(funcs).ParseGlob("./generate/templates/*")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	data := map[string]interface{}{
		"Builtins": builtins,
	}
	for _, t := range templates.Templates() {
		// template with Name = "foo-bar.go.template" will be written to ./foo/bar.go
		// "bar.go.template" ./bar.go
		// "fizz-buzz-bar.go.template" ./fizz/buzz/bar.go

		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("code generation failed: %v", err)
		}

		fullname := filepath.Join(pwd, templateFilePath(t.Name()))
		f, err := os.Create(fullname)
		if err != nil {
			log.Fatalf("failed to create %v: %v", fullname, err)
		}
		if err := t.Execute(f, data); err != nil {
			log.Fatalf("failed to execute %v: %v", t.Name(), err)
		}
		f.Close()

		gofmt := exec.Command("gofmt", "-s", "-w", fullname)
		gofmt.Stdout = os.Stdout
		gofmt.Stderr = os.Stderr
		if err := gofmt.Run(); err != nil {
			log.Fatalf("failed to run gofmt on %v: %v", fullname, err)
		}
	}
}

// Determines the relative file path for a templated file
// given the name of the template.
//
// Dashes in template names are converted to directory separators
// and the .template suffix is removed.
//
// For example:
//
//	foo-bar.go.template      => foo/bar.go
//	bar.go.template          => bar.go
//	fizz-buz-bar.go.template => fizz/buz/bar.go
func templateFilePath(name string) string {
	parts := strings.Split(name, "-")
	filename := strings.TrimSuffix(parts[len(parts)-1], ".template")
	parts[len(parts)-1] = filename
	return filepath.Join(parts...)
}

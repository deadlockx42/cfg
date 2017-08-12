//
//   Copyright 2017 Deadlock X42 <deadlock.x42@gmail.com>
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//

package schema

import (
	"fmt"
	"strings"
	"unicode"
)

// ValidationResults is a list of warnings and errors that are produced by validating
// a generator.
type ValidationResults struct {
	Warnings []string
	Errors   []string
}

// Validate checks the correctness of the generator structure.
//
// For a generator it will
// 	- warn if the copyright is empty
//  - error if the "Begin" type isn't defined
//
// For an object it will
//  - warn if the documentation is empty
//	- error if the name of the object is not a valid identifier
//	- error if the name of the object is a reserved or predeclared word
//
// For each field of an object it will
//	- error if the name of the field is not a valid identifier
//	- error if the type of the field is not defined
//
// For fields with precludes defined it will
//	- TODO: error if the a field name in precludes does not exist
//
// For an array it will
//  - warn if the documentation is empty
//	- error if the name of the array is not a valid identifier
//	- error if the type of the array is not defined
//
func Validate(g Generator) (*ValidationResults, error) {
	r := &ValidationResults{
		Warnings: []string{},
		Errors:   []string{},
	}
	err := g.Accept(r)
	return r, err
}

// VisitGenerator validates a generator.
func (r *ValidationResults) VisitGenerator(g Generator) error {
	if len(g.Copyright()) == 0 {
		r.Warnings = append(r.Warnings, "Empty copyright")
	}
	if !isType(g.Begin()) {
		r.Errors = append(r.Errors, fmt.Sprintf("Begin %q not defined.", g.Begin()))
	}
	return nil
}

// VisitObject validates a generator object.
func (r *ValidationResults) VisitObject(o Object) error {
	if !isName(o.Name()) {
		r.Errors = append(r.Errors, fmt.Sprintf("Object name %q is not a valid identifier.", o.Name()))
	}
	if len(o.Documentation()) == 0 {
		r.Warnings = append(r.Warnings, fmt.Sprintf("Documentation for object %q does not exist.", o.Name()))
	}
	for _, f := range o.Fields() {
		if !isIdentifier(f.Name()) {
			r.Errors = append(r.Errors, fmt.Sprintf("Object %q, field name %q is not a valid identifier.", o.Name(), f.Name()))
		}
		if !isType(f.Type()) {
			r.Errors = append(r.Errors, fmt.Sprintf("Object %q, field name %q, type %q not defined.", o.Name(), f.Name(), f.Type()))
		}
		for _, p := range f.Precludes() {
			found := false
			for _, f := range o.Fields() {
				if p == f.Name() {
					found = true
					break
				}
			}
			if !found {
				r.Errors = append(r.Errors, fmt.Sprintf("Object %q, field name %q, precludes an unknown field %q.", o.Name(), f.Name(), p))
			}
		}
	}
	return nil
}

// VisitArray validates a generator array.
func (r *ValidationResults) VisitArray(a Array) error {
	if !isIdentifier(a.Name()) {
		r.Errors = append(r.Errors, fmt.Sprintf("Array name %q is not a valid identifier.", a.Name()))
	}
	if len(a.Documentation()) == 0 {
		r.Warnings = append(r.Warnings, fmt.Sprintf("Documentation for array %q does not exist.", a.Name()))
	}
	if !isType(a.Type()) {
		r.Errors = append(r.Errors, fmt.Sprintf("Array type %q not defined.", a.Type()))
	}
	return nil
}

var reserved = map[string]bool{
	"break": true, "default": true, "func": true, "interface": true, "select": true,
	"case": true, "defer": true, "go": true, "map": true, "struct": true,
	"chan": true, "else": true, "goto": true, "package": true, "switch": true,
	"const": true, "fallthrough": true, "if": true, "range": true, "type": true,
	"continue": true, "for": true, "import": true, "return": true, "var": true,
}

func isReserved(s string) bool {
	return reserved[s]
}

var predeclared = map[string]bool{
	"bool": true, "byte": true, "complex64": true, "complex128": true, "error": true,
	"float32": true, "float64": true, "int": true, "int8": true, "int16": true,
	"int32": true, "int64": true, "rune": true, "string": true, "uint": true,
	"uint8": true, "uint16": true, "uint32": true, "uint64": true, "uintptr": true,
	"true": true, "false": true, "iota": true, "nil": true, "append": true,
	"cap": true, "close": true, "complex": true, "copy": true, "delete": true,
	"imag": true, "len": true, "make": true, "new": true, "panic": true,
	"print": true, "println": true, "real": true, "recover": true,
}

func isPredeclared(s string) bool {
	return predeclared[s]
}

func isIdentifier(name string) bool {
	s := strings.ToLower(name)
	first := true
	for _, r := range s {
		switch first {
		case true:
			first = false
			if unicode.IsLetter(r) || r == '_' {
				continue
			}
			return false
		case false:
			if unicode.IsLetter(r) || r == '_' || unicode.IsDigit(r) {
				continue
			}
			return false
		}
	}
	return true
}

var types = map[string]bool{
	"bool": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
	"int8": true, "int16": true, "int32": true, "int64": true, "float32": true,
	"float64": true, "complex64": true, "complex128": true, "byte": true, "rune": true,
	"uint": true, "int": true, "uintptr": true, "string": true,
}

func isType(s string) bool {
	if maps.objects[s] != nil || maps.arrays[s] != nil {
		return true
	}
	return types[s]
}

func isName(s string) bool {
	name := strings.ToLower(s)
	if !isIdentifier(name) || isReserved(name) || isPredeclared(name) {
		return false
	}
	return true
}

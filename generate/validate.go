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

package generate

import "fmt"

type ValidationError struct {
	string
}

func (i ValidationError) Error() string {
	return i.string
}

type ValidationResults struct {
	Warnings []string
	Errors   []string
}

type inputValidator struct {
	maps    *Maps
	results *ValidationResults
}

func Validate(g Generator, m *Maps) (*ValidationResults, error) {
	if m == nil {
		return nil, ValidationError{"Validate: nil maps"}
	}
	i := &inputValidator{
		maps: m,
		results: &ValidationResults{
			Warnings: []string{},
			Errors:   []string{},
		},
	}
	err := g.Accept(i)
	return i.results, err
}

func (i *inputValidator) VisitGenerator(g Generator) error {
	// Warn if the copyright isn't defined.
	if len(g.Copyright()) == 0 {
		i.results.Warnings = append(i.results.Warnings, "Empty copyright")
	}
	// Error if the begin type isn't defined.
	if i.maps.Objects[g.Begin()] == nil && i.maps.Arrays[g.Begin()] == nil {
		e := fmt.Sprintf("Begin type %s not defined", g.Begin())
		i.results.Errors = append(i.results.Errors, e)
	}
	return nil
}

func (i *inputValidator) VisitObject(o Object) error {
	// TODO
	return nil
}

func (i *inputValidator) VisitArray(a Array) error {
	// TODO
	return nil
}

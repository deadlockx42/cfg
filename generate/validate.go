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

type ValidationResults struct {
	Warnings []string
	Errors   []string
}

type inputValidator struct {
	results *ValidationResults
}

func Validate(g Generator) (*ValidationResults, error) {
	r := &ValidationResults{
		Warnings: []string{},
		Errors:   []string{},
	}
	err := g.Accept(r)
	return r, err
}

func (r *ValidationResults) VisitGenerator(g Generator) error {
	// Warn if the copyright isn't defined.
	if len(g.Copyright()) == 0 {
		r.Warnings = append(r.Warnings, "Empty copyright")
	}
	// Error if the begin type isn't defined.
	if maps.objects[g.Begin()] == nil && maps.arrays[g.Begin()] == nil {
		e := fmt.Sprintf("Begin type %s not defined", g.Begin())
		r.Errors = append(r.Errors, e)
	}
	return nil
}

func (r *ValidationResults) VisitObject(o Object) error {
	// TODO
	return nil
}

func (r *ValidationResults) VisitArray(a Array) error {
	// TODO
	return nil
}

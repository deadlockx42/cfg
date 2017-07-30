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

type Maps struct {
	Objects map[string]Object
	Arrays  map[string]Array
}

func NewMaps(g Generator) (*Maps, error) {
	m := &Maps{
		Objects: map[string]Object{},
		Arrays:  map[string]Array{},
	}
	if err := g.Accept(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (*Maps) VisitGenerator(Generator) error {
	return nil
}

func (m *Maps) VisitObject(o Object) error {
	m.Objects[o.Name()] = o
	return nil
}

func (m *Maps) VisitArray(a Array) error {
	m.Arrays[a.Name()] = a
	return nil
}

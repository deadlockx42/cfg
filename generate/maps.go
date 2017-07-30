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

type genmaps struct {
	objects map[string]Object
	arrays  map[string]Array
}

var maps *genmaps

func init() {
	initializers = append(initializers, newMaps)
}

func newMaps(g Generator) error {
	maps = &genmaps{
		objects: map[string]Object{},
		arrays:  map[string]Array{},
	}
	if err := g.Accept(maps); err != nil {
		return err
	}
	return nil
}

func (*genmaps) VisitGenerator(Generator) error {
	return nil
}

func (m *genmaps) VisitObject(o Object) error {
	m.objects[o.Name()] = o
	return nil
}

func (m *genmaps) VisitArray(a Array) error {
	m.arrays[a.Name()] = a
	return nil
}

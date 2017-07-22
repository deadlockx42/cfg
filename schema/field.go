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

// Field represents an element of an object. It contains a name, type, tag and precludes.
type Field interface {
	Name() string
	Type() string
	Tag() string
	Precludes() Precludes
}

type field struct {
	FieldName      string    `json:"Field"`
	FieldType      string    `json:"Type"`
	FieldTag       string    `json:"Tag"`
	FieldPrecludes Precludes `json:"Precludes"`
}

func (f *field) Name() string {
	return f.FieldName
}

func (f *field) Type() string {
	return f.FieldType
}

func (f *field) Tag() string {
	return f.FieldTag
}

func (f *field) Precludes() Precludes {
	return f.FieldPrecludes
}

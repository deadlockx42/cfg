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

// Object contains a name and a set of fields.
type Object interface {
	Acceptor
	Name() string
	Documentation() Text
	Fields() Fields
	Acceptor() bool
}

type object struct {
	OName          string `json:"Object"`
	ODocumentation Text   `json:"Documentation"`
	OFields        Fields `json:"Fields"`
	OAcceptor      bool   `json:"Acceptor"`
}

func (o *object) Accept(v Visitor) error {
	return v.VisitObject(o)
}

func (o *object) Name() string {
	return o.OName
}

func (o *object) Documentation() Text {
	return o.ODocumentation
}

func (o *object) Fields() Fields {
	return o.OFields
}

func (o *object) Acceptor() bool {
	return o.OAcceptor
}

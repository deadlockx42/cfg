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

// Array has a name and a type.
type Array interface {
	Acceptor
	Name() string
	Type() string
	Documentation() Text
}

type array struct {
	AName          string `json:"Array"`
	AType          string `json:"Type"`
	ADocumentation Text   `json:"Documentation"`
}

func (a *array) Accept(v Visitor) error {
	return v.VisitArray(a)
}

func (a *array) Name() string {
	return a.AName
}

func (a *array) Type() string {
	return a.AType
}

func (a *array) Documentation() Text {
	return a.ADocumentation
}

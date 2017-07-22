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

// Array has a name and a type.
type Array interface {
	Name() string
	Type() string
}

type array struct {
	ArrayName string `json:"Array"`
	ArrayType string `json:"Type"`
}

func (a *array) Name() string {
	return a.ArrayName
}

func (a *array) Type() string {
	return a.ArrayType
}

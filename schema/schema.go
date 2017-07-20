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

type Schema interface {
	Copyright() Copyright
	Begin() string
}

type schema struct {
	Copyright_   Copyright    `json:"Copyright"`
	Begin_       string       `json:"Begin"`
	Definitions_ *definitions `json:"Definitions"`
}

func (s *schema) Copyright() Copyright {
	return s.Copyright_
}

func (s *schema) Begin() string {
	return s.Begin_
}

func (s *schema) Definitions() Definitions {
	return s.Definitions_
}

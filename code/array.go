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

package code

import (
	// "fmt"
	"io"
	// "strings"

	"github.com/deadlockx42/voidgen/schema"
)

type arraySrc struct {
	*source
	schema schema.Generator
	array  schema.Array
}

func Array(pkg string, s schema.Generator, a schema.Array) *arraySrc {
	return &arraySrc{
		source: &source{
			copyright: s.Copyright(),
			pkg:       pkg,
		},
		schema: s,
		array:  a,
	}
}

func (s *arraySrc) Write(w io.Writer) (int, error) {
	n, err := s.source.Write(w)
	if err != nil {
		return n, err
	}

	b := []byte{}
	// TODO
	return w.Write(b)
}

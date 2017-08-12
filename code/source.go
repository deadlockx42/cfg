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
	"io"
	"strings"

	"github.com/deadlockx42/voidgen/schema"
)

type source struct {
	copyright []string
	pkg       string
}

func newSource(pkg string, s schema.Generator) *source {
	return &source{
		copyright: s.Copyright(),
		pkg:       pkg,
	}
}

func (s *source) Write(w io.Writer) (n int, err error) {
	b := ""
	for _, c := range s.copyright {
		b += strings.TrimSpace("//   "+c) + "\n"
	}
	b += "\npackage " + s.pkg + "\n"
	return w.Write([]byte(b))
}

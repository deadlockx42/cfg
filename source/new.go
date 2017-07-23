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

package source

import (
	"fmt"
	"io"
	"strings"

	"github.com/deadlockx42/voidgen/schema"
)

type newSrc struct {
	*source
	schema schema.Schema
}

func New(pkg string, s schema.Schema) *newSrc {
	return &newSrc{
		source: &source{
			copyright: s.Copyright(),
			pkg:       pkg,
		},
		schema: s,
	}
}

func (s *newSrc) Write(w io.Writer) (int, error) {
	n, err := s.source.Write(w)
	if err != nil {
		return n, err
	}

	b := []byte{}

	imports := "\nimport (\n" +
		"\t\"encoding/json\"\n" +
		"\t\"io\"\n" +
		")\n\n"
	b = append(b, []byte(imports)...)

	newfunc := "// %s\n" +
		"func New(r io.Reader) (%s, error) {\n" +
		"\tvar %s %s\n" +
		"\tfor {\n" +
		"\t\terr := json.NewDecoder(r).Decode(&%s)\n" +
		"\t\tif err == io.EOF {\n" +
		"\t\t\tbreak\n" +
		"\t\t}\n" +
		"\t\tif err != nil {\n" +
		"\t\t\treturn nil, err\n" +
		"\t\t}\n" +
		"\t}\n" +
		"\treturn &%s, nil\n" +
		"}\n"

	Begin := s.schema.Begin()
	begin := strings.ToLower(Begin[0:1]) + Begin[1:]
	c := begin[0:1]
	comment := "New creates a " + begin + "."
	b = append(b, []byte(fmt.Sprintf(newfunc, comment, Begin, c, begin, c, c))...)
	return w.Write(b)
}

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
	// "os"
	// "path/filepath"
	"strings"
	"unicode"
	// "github.com/deadlockx42/voidgen/schema"
)

/*
func renderDoc(s schema.Generator, output, pkg string) error {
	file, err := os.OpenFile(filename(output, "doc"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = source.Doc(pkg, s).Write(file)
	if err != nil {
		file.Close()
		return err
	}
	return nil
}

func renderNew(s schema.Generator, output, pkg string) error {
	file, err := os.OpenFile(filename(output, "new"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = source.New(pkg, s).Write(file)
	if err != nil {
		return err
	}
	return nil
}

func renderObjects(s schema.Generator, output, pkg string) error {
	for _, o := range s.Definitions().Objects() {
		file, err := os.OpenFile(filename(output, o.Name()), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = source.Object(pkg, s, o).Write(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func renderArrays(s schema.Generator, output, pkg string) error {
	for _, a := range s.Definitions().Arrays() {
		file, err := os.OpenFile(filename(output, a.Name()), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = source.Array(pkg, s, a).Write(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func render(s schema.Generator, output string) error {
	if err := os.MkdirAll(output, 0755); err != nil {
		return err
	}
	pkg := filepath.Base(output)

	if err := renderDoc(s, output, pkg); err != nil {
		return err
	}
	if err := renderNew(s, output, pkg); err != nil {
		return err
	}
	if err := renderObjects(s, output, pkg); err != nil {
		return err
	}
	if err := renderArrays(s, output, pkg); err != nil {
		return err
	}
	return nil
}
*/

func filename(dir, name string) string {
	f := strings.TrimSpace(name)
	f = dir + "/" + strings.Replace(f, " ", "_", -1)
	f += ".go"

	s := []rune{}
	var prev rune
	first := true

	for _, c := range f {
		if unicode.IsUpper(c) {
			if first {
				first = false
			} else {
				if prev != '_' {
					s = append(s, rune('_'))
				}
			}
			s = append(s, unicode.ToLower(c))
		} else {
			s = append(s, c)
		}
		prev = c
	}
	return string(s)
}

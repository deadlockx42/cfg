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

import (
	"encoding/json"
	"io"
)

// Initializiers allow for additional functionality to be created as part of the
// generator construction. To take advantage of this, append an initializer
// function to the initializers slice in an init() function.
type initializer func(Generator) error

var initializers []initializer

// New creates a generator.
func New(r io.Reader) (Generator, error) {
	g := &generator{}
	for {
		err := json.NewDecoder(r).Decode(g)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	for _, i := range initializers {
		if err := i(g); err != nil {
			return nil, err
		}
	}
	return g, nil
}

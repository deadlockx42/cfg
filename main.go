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

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/deadlockx42/voidgen/generate"
)

var (
	verbose bool
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s file\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	g, err := generate.New(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	m, err := generate.NewMaps(g)
	if err != nil {
		log.Fatal(err.Error())
	}

	results, err := generate.Validate(g, m)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, w := range results.Warnings {
		fmt.Printf("Warning: %s\n", w)
	}
	for _, e := range results.Errors {
		fmt.Printf("Error: %s\n", e)
	}
}

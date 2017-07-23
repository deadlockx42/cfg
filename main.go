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
	"flag"
	"log"
	"os"

	"github.com/deadlockx42/voidgen/schema"
)

func main() {
	var file, output string
	flag.StringVar(&file, "file", "", "input scheman file")
	flag.StringVar(&file, "f", "", "input scheman file")
	flag.StringVar(&output, "output", "", "output directory")
	flag.StringVar(&output, "o", "", "output directory")
	flag.Parse()

	if output == "" {
		log.Fatal("Output directory required.")
	}
	if output == "." || output == ".." {
		log.Fatalf("Output directory %q not allowed.", output)
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	s, err := schema.New(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = render(s, output)
	if err != nil {
		log.Fatal(err.Error())
	}
}

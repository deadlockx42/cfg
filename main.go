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

	"github.com/deadlockx42/voidgen/generate"
	"github.com/deadlockx42/voidgen/schema"
)

func main() {
	var file, output string
	flag.StringVar(&file, "file", "", "input scheman file")
	flag.StringVar(&file, "f", "", "input scheman file")
	flag.StringVar(&output, "output", ".", "output directory")
	flag.StringVar(&output, "o", ".", "output directory")
	flag.Parse()

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	if output != "." {
		if err := os.Mkdir(output, 0644); err != nil {
			log.Fatal(err.Error())
		}
	}

	s, err := schema.New(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	g, err := generate.New(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = g.Write(output)
	if err != nil {
		log.Fatal(err.Error())
	}
}

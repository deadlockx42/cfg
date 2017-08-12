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

// Generator generates go code based on a json definition file.
type Generator interface {
	Acceptor
	Name() string
	Copyright() Text
	Begin() string
	Objects() Objects
	Arrays() Arrays
}

type generator struct {
	GName      string  `json:"Name"`
	GCopyright Text    `json:"Copyright"`
	GBegin     string  `json:"Begin"`
	GObjects   Objects `json:"Objects"`
	GArrays    Arrays  `json:"Arrays"`
}

func (g *generator) Accept(v Visitor) error {
	if err := v.VisitGenerator(g); err != nil {
		return err
	}
	if err := g.Objects().Accept(v); err != nil {
		return err
	}
	if err := g.Arrays().Accept(v); err != nil {
		return err
	}
	return nil
}

func (g *generator) Name() string {
	return g.GName
}

func (g *generator) Copyright() Text {
	return g.GCopyright
}

func (g *generator) Begin() string {
	return g.GBegin
}

func (g *generator) Objects() Objects {
	return g.GObjects
}

func (g *generator) Arrays() Arrays {
	return g.GArrays
}

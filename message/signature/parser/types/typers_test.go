// Copyright © 2017 Dmitry Moskowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package types

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var typersSamples = []struct {
	t  Typer
	s  string
	k  Kind
	ts Types
}{
	{
		Typers{
			String,
			Int32,
		},
		"[ String, Int32 ]",
		Complex,
		Types{String, Int32},
	},
	{
		Typers{
			&ComplexType{
				Typer: Array,
				Body:  String,
			},
			Int32,
		},
		"[ Array String, Int32 ]",
		Complex,
		Types{String, Int32},
	},
	{
		Typers{},
		"[ ]",
		Complex,
		Types{},
	},
	{
		(Typers)(nil),
		"[ ]",
		Complex,
		Types{},
	},
	{
		Typers{Type(0)},
		"[ Invalid ]",
		Complex,
		Types{Type(0)},
	},
	{
		Typers{Type(255), Type(255)},
		"[ Invalid, Invalid ]",
		Complex,
		Types{Type(255), Type(255)},
	},
}

func TestTypersString(t *testing.T) {
	for k, sample := range typersSamples {
		assert.Equal(
			t,
			sample.s,
			sample.t.String(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTypersKind(t *testing.T) {
	for k, sample := range typersSamples {
		assert.Equal(
			t,
			sample.k,
			sample.t.Kind(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTypersTypes(t *testing.T) {
	for k, sample := range typersSamples {
		assert.Equal(
			t,
			sample.ts,
			sample.t.Types(),
			spew.Sdump(k, sample),
		)
	}
}

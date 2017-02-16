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

var simpleSamples = []struct {
	t  Typer
	s  string
	k  Kind
	ts Types
}{
	{
		&SimpleType{Byte},
		"Byte",
		Simple,
		Types{Byte},
	},
	{
		&SimpleType{},
		"",
		Simple,
		Types{},
	},
	{
		&SimpleType{
			&ComplexType{
				Typer: Array,
				Body:  String,
			},
		},
		"Array String",
		Simple,
		Types{String},
	},
	{
		&SimpleType{Type(0)},
		"Invalid",
		Simple,
		Types{Type(0)},
	},
	{
		&SimpleType{Type(255)},
		"Invalid",
		Simple,
		Types{Type(255)},
	},
}

func TestSimpleString(t *testing.T) {
	for k, sample := range simpleSamples {
		assert.Equal(
			t,
			sample.s,
			sample.t.String(),
			spew.Sdump(k, sample),
		)
	}
}

func TestSimpleKind(t *testing.T) {
	for k, sample := range simpleSamples {
		assert.Equal(
			t,
			sample.k,
			sample.t.Kind(),
			spew.Sdump(k, sample),
		)
	}
}

func TestSimpleTypes(t *testing.T) {
	for k, sample := range simpleSamples {
		assert.Equal(
			t,
			sample.ts,
			sample.t.Types(),
			spew.Sdump(k, sample),
		)
	}
}
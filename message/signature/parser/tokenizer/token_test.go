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
package tokenizer

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/corpix/dbus/message/signature/parser/types"
)

var tokenSamples = []struct {
	r  Tokener
	s  string
	ts Tokens
	e  string
	t  types.Types
}{
	{
		Int64,
		"Int64",
		Tokens{Int64},
		"x",
		types.Types{types.Int64},
	},
	{
		Array,
		"Array",
		Tokens{Array},
		"a",
		types.Types{types.Array},
	},
	{
		Struct,
		"Struct",
		Tokens{Struct},
		"",
		types.Types{types.Struct},
	},
	{
		Dict,
		"Dict",
		Tokens{Dict},
		"",
		types.Types{types.Dict},
	},
	{
		Token(255),
		"Invalid",
		Tokens{Token(255)},
		"",
		types.Types{types.Invalid},
	},
	{
		Token(0),
		"Invalid",
		Tokens{Invalid},
		"",
		types.Types{types.Invalid},
	},
}

func TestTokenString(t *testing.T) {
	for k, sample := range tokenSamples {
		assert.Equal(
			t,
			sample.s,
			sample.r.String(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTokenTokens(t *testing.T) {
	for k, sample := range tokenSamples {
		assert.Equal(
			t,
			sample.ts,
			sample.r.Tokens(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTokenEncode(t *testing.T) {
	for k, sample := range tokenSamples {
		assert.Equal(
			t,
			sample.e,
			sample.r.Encode(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTokenTypes(t *testing.T) {
	for k, sample := range tokenSamples {
		assert.Equal(
			t,
			sample.t,
			sample.r.Types(),
			spew.Sdump(k, sample),
		)
	}
}

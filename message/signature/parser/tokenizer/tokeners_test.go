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

var tokenersSamples = []struct {
	r  Tokener
	s  string
	ts Tokens
	e  string
	t  types.Types
}{
	{
		Tokeners{Int64, Int16},
		"[ Int64, Int16 ]",
		Tokens{Int64, Int16},
		"xn",
		types.Types{types.Int64, types.Int16},
	},
	{
		Tokeners{Array},
		"[ Array ]",
		Tokens{Array},
		"a",
		types.Types{types.Array},
	},
	{
		Tokeners{Struct, StructStart, StructEnd},
		"[ Struct, StructStart, StructEnd ]",
		Tokens{Struct, StructStart, StructEnd},
		"()",
		types.Types{types.Struct, types.Invalid, types.Invalid},
	},
	{
		Tokeners{Dict, DictStart, DictEnd},
		"[ Dict, DictStart, DictEnd ]",
		Tokens{Dict, DictStart, DictEnd},
		"{}",
		types.Types{types.Dict, types.Invalid, types.Invalid},
	},
	{
		Tokeners{Token(255), Int16},
		"[ Invalid, Int16 ]",
		Tokens{Token(255), Int16},
		"n",
		types.Types{types.Invalid, types.Int16},
	},
	{
		Tokeners{Token(0)},
		"[ Invalid ]",
		Tokens{Invalid},
		"",
		types.Types{types.Invalid},
	},
	{
		Tokeners{
			RecursiveToken{
				Tokener: Array,
				Body:    Tokens{Int16, Int32},
				Head:    DictStart,
				Tail:    DictEnd,
			},
			String,
		},
		"[ Array DictStart [ Int16, Int32 ] DictEnd, String ]",
		Tokens{Array, DictStart, Int16, Int32, DictEnd, String},
		"a{ni}s",
		types.Types{types.Array, types.Int16, types.Int32, types.String},
	},
}

func TestTokenersString(t *testing.T) {
	for k, sample := range tokenersSamples {
		assert.Equal(
			t,
			sample.s,
			sample.r.String(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTokenersTokens(t *testing.T) {
	for k, sample := range tokenersSamples {
		assert.Equal(
			t,
			sample.ts,
			sample.r.Tokens(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTokenersEncode(t *testing.T) {
	for k, sample := range tokenersSamples {
		assert.Equal(
			t,
			sample.e,
			sample.r.Encode(),
			spew.Sdump(k, sample),
		)
	}
}

func TestTokenersTypes(t *testing.T) {
	for k, sample := range tokenersSamples {
		assert.Equal(
			t,
			sample.t,
			sample.r.Types(),
			spew.Sdump(k, sample),
		)
	}
}

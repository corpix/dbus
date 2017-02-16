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
package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/corpix/dbus/message/signature/parser/tokenizer"
	"github.com/corpix/dbus/message/signature/parser/types"
)

func TestParse(t *testing.T) {
	samples := []struct {
		signature      string
		representation string
		kind           types.Kind
		err            error
	}{
		{
			"y",
			"[ Byte ]",
			types.Complex,
			nil,
		},
		{
			"(ii)",
			"[ Struct [ Int32, Int32 ] ]",
			types.Complex,
			nil,
		},
		{
			"(ii({x}{x}))",
			"[ Struct [ Int32, Int32, Struct [ Dict [ Int64 ], Dict [ Int64 ] ] ] ]",
			types.Complex,
			nil,
		},
		{
			"xa{iia{s}}iiix",
			"[ Int64, Array [ Int32, Int32, Array [ String ] ], Int32, Int32, Int32, Int64 ]",
			types.Complex,
			nil,
		},
		{
			"r(ii)",
			"",
			types.Complex,
			tokenizer.NewErrVirtualToken(1, tokenizer.Struct),
		},
	}

	for _, sample := range samples {
		tr, err := DefaultParser.Parse(sample.signature)
		assert.Equal(
			t,
			sample.err,
			err,
		)
		if sample.err == nil {
			assert.Equal(
				t,
				sample.representation,
				tr.String(),
			)
			assert.Equal(
				t,
				sample.kind.String(),
				tr.Kind().String(),
			)
		} else {
			assert.Equal(t, nil, tr)
		}
	}
}

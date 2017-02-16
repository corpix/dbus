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
)

func TestBoundsGetBounds(t *testing.T) {
	samples := []struct {
		pos    int
		ts     Tokens
		start  Token
		end    Token
		result *Bounds
		err    error
	}{
		{
			0,
			Tokens{},
			DictStart,
			DictEnd,
			&Bounds{0, 0},
			nil,
		},
		{
			0,
			Tokens{
				Array,
				DictStart,
				String,
				DictEnd,
			},
			DictStart,
			DictEnd,
			&Bounds{1, 3},
			nil,
		},
		{
			0,
			Tokens{
				Array,
				DictStart,
				DictEnd,
			},
			DictStart,
			DictEnd,
			&Bounds{1, 2},
			nil,
		},
		{
			0,
			Tokens{
				Array,
				DictStart,
				Array,
				DictStart,
				String,
				DictEnd,
				DictEnd,
			},
			DictStart,
			DictEnd,
			&Bounds{1, 6},
			nil,
		},
		{
			0,
			Tokens{
				Array,
				DictStart,
				String,
			},
			DictStart,
			DictEnd,
			nil,
			NewErrBoundIncomplete(2, DictStart, DictEnd),
		},
	}

	var (
		bounds *Bounds
		err    error
	)

	for k, sample := range samples {
		bounds, err = GetBounds(
			sample.pos,
			sample.ts,
			sample.start,
			sample.end,
		)
		assert.Equal(
			t,
			sample.err,
			err,
			spew.Sdump(k, sample),
		)
		if sample.err == nil {
			assert.Equal(
				t,
				sample.result,
				bounds,
				spew.Sdump(k, sample),
			)
		} else {
			assert.Equal(
				t,
				(*Bounds)(nil),
				bounds,
				spew.Sdump(k, sample),
			)
		}
	}
}

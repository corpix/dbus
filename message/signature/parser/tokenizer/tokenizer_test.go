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

	"github.com/corpix/dbus/message/signature/parser/errors"
)

func TestTokenizerTokenize(t *testing.T) {
	nesting := 10
	samples := []struct {
		s   string
		t   Tokens
		err error
	}{
		{
			"",
			Tokens{},
			nil,
		},
		{
			"y",
			Tokens{Byte},
			nil,
		},
		{
			"yi",
			Tokens{Byte, Int32},
			nil,
		},
		{
			"a{i}",
			Tokens{Array, DictStart, Int32, DictEnd},
			nil,
		},
		{
			"xa{x}",
			Tokens{Int64, Array, DictStart, Int64, DictEnd},
			nil,
		},
		{
			"qutdsoa{(gh)v}",
			Tokens{
				Uint16,
				Uint32,
				Uint64,
				Double,
				String,
				ObjectPath,
				Array,
				DictStart,
				Struct,
				StructStart,
				Signature,
				UnixFD,
				StructEnd,
				Variant,
				DictEnd,
			},
			nil,
		},
		{
			"a{(xx)(x)}",
			Tokens{
				Array,
				DictStart,
				Struct,
				StructStart,
				Int64,
				Int64,
				StructEnd,
				Struct,
				StructStart,
				Int64,
				StructEnd,
				DictEnd,
			},
			nil,
		},
		{
			"{a{s}{(i)(i(x(u)))}}",
			Tokens{
				Dict,
				DictStart,
				Array,
				DictStart,
				String,
				DictEnd,
				Dict,
				DictStart,
				Struct,
				StructStart,
				Int32,
				StructEnd,
				Struct,
				StructStart,
				Int32,
				Struct,
				StructStart,
				Int64,
				Struct,
				StructStart,
				Uint32,
				StructEnd,
				StructEnd,
				StructEnd,
				DictEnd,
				DictEnd,
			},
			nil,
		},
		{
			"a{a{a{a{a{a{a{a{a{a{a{a{}}}}}}}}}}}}",
			Tokens{},
			errors.NewErrNestingTooDeep(21, nesting),
		},
		{
			"r(i)",
			Tokens{},
			NewErrVirtualToken(1, Struct),
		},
		{
			"xixixixir(i)",
			Tokens{},
			NewErrVirtualToken(9, Struct),
		},
		{
			"xixixixi(((r(i))))",
			Tokens{},
			NewErrVirtualToken(12, Struct),
		},
		{
			"e{i}",
			Tokens{},
			NewErrVirtualToken(1, Dict),
		},
		{
			"xixixixie{i}",
			Tokens{},
			NewErrVirtualToken(9, Dict),
		},
		{
			"xixixixi(((e{i})))",
			Tokens{},
			NewErrVirtualToken(12, Dict),
		},
		{
			"ixijxixi(((e{i})))",
			Tokens{},
			NewErrUnsupportedToken(4, 'j'),
		},
	}

	tokenizer := NewTokenizer(nesting)
	var (
		tr  Tokener
		err error
	)

	for _, sample := range samples {
		tr, err = tokenizer.Tokenize(sample.s)
		assert.Equal(
			t,
			sample.err,
			err,
			spew.Sdump(sample),
		)
		if sample.err == nil {
			assert.Equal(
				t,
				sample.t,
				tr.Tokens(),
				spew.Sdump(sample),
			)
		} else {
			assert.Equal(
				t,
				nil,
				tr,
				spew.Sdump(sample),
			)
		}
	}
}

func TestTokenizerFold(t *testing.T) {
	samples := []struct {
		ts     Tokens
		result Tokens
		err    error
	}{
		{
			Tokens{Int32, Token(255)},
			Tokens{},
			NewErrUnsupportedToken(2, 255),
		},
	}

	for _, sample := range samples {
		tr, err := NewTokenizer(10).Fold(sample.ts)
		assert.Equal(
			t,
			sample.err,
			err,
			spew.Sdump(sample),
		)
		if sample.err == nil {
			assert.Equal(
				t,
				sample.result,
				tr.Tokens(),
				spew.Sdump(sample),
			)
		} else {
			assert.Equal(
				t,
				nil,
				tr,
				spew.Sdump(sample),
			)
		}
	}
}

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
	"strings"

	"github.com/corpix/dbus/message/signature/parser/types"
)

type RecursiveToken struct {
	Tokener
	Body Tokener
	Head Tokener
	Tail Tokener
}

func (t RecursiveToken) String() string {
	ts := []string{}
	if t.Tokener != nil {
		ts = append(ts, t.Tokener.String())
	}
	if t.Head != nil {
		ts = append(ts, t.Head.String())
	}
	if t.Body != nil {
		ts = append(ts, t.Body.String())
	}
	if t.Tail != nil {
		ts = append(ts, t.Tail.String())
	}

	return strings.Join(
		ts,
		" ",
	)
}

func (t RecursiveToken) Tokens() Tokens {
	ts := Tokens{}

	if t.Tokener != nil {
		ts = append(ts, t.Tokener.Tokens()...)
	}
	if t.Head != nil {
		ts = append(ts, t.Head.Tokens()...)
	}
	if t.Body != nil {
		ts = append(ts, t.Body.Tokens()...)
	}
	if t.Tail != nil {
		ts = append(ts, t.Tail.Tokens()...)
	}

	return ts
}

func (t RecursiveToken) Encode() string {
	ts := ""

	for _, v := range t.Tokens() {
		ts += v.Encode()
	}

	return ts
}

func (t RecursiveToken) Types() types.Types {
	ts := types.Types{}

	if t.Tokener != nil {
		ts = append(ts, t.Tokener.Types()...)
	}
	if t.Body != nil {
		ts = append(ts, t.Body.Types()...)
	}

	return ts
}

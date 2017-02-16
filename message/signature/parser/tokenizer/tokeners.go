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
	"fmt"
	"strings"

	"github.com/corpix/dbus/message/signature/parser/types"
)

type Tokeners []Tokener

func (t Tokeners) String() string {
	tokenStrings := make([]string, len(t))
	for k, v := range t {
		tokenStrings[k] = v.String()
	}

	return fmt.Sprintf(
		"[ %s ]",
		strings.Join(tokenStrings, ", "),
	)
}

func (t Tokeners) Tokens() Tokens {
	ts := Tokens{}

	for _, v := range t {
		ts = append(ts, v.Tokens()...)
	}

	return ts
}

func (t Tokeners) Encode() string {
	encoded := ""

	for _, v := range t {
		encoded += v.Encode()
	}

	return encoded
}

func (t Tokeners) Types() types.Types {
	ts := types.Types{}

	for _, v := range t {
		ts = append(ts, v.Types()...)
	}

	return ts
}

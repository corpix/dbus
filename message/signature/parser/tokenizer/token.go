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
	"github.com/corpix/dbus/message/signature/parser/types"
)

type Token uint

func (t Token) String() string {
	if v, ok := TokenNames[t]; ok {
		return v
	}

	return TokenNames[Invalid]
}

func (t Token) Tokens() Tokens {
	return Tokens{t}
}

func (t Token) Encode() string {
	// Virtual tokens has no string representation
	// see types.go#TokenVirtual comment.
	if _, ok := TokenVirtual[t]; ok {
		return ""
	}
	if v, ok := TokenEncoded[t]; ok {
		return string(v)
	}

	return ""
}

func (t Token) Types() types.Types {
	if v, ok := TokenType[t]; ok {
		return types.Types{v}
	}

	return types.Types{types.Invalid}
}

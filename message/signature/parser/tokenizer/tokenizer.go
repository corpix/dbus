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
	"github.com/corpix/dbus/message/signature/parser/errors"
)

var (
	DefaultTokenizer = &Tokenizer{maxDepth: 64}
)

type Tokenizer struct {
	maxDepth int
}

func (t *Tokenizer) Tokenize(s string) (Tokener, error) {
	tokens := make(Tokens, len(s))
	for k, v := range s {
		if t, ok := EncodedToken[v]; ok {
			tokens[k] = t
		} else {
			return nil, NewErrUnsupportedToken(k+1, int(v))
		}
	}

	return t.Fold(tokens)
}

func (t *Tokenizer) Fold(tokens Tokener) (Tokener, error) {
	ts, err := t.fold(1, 0, tokens.Tokens())
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (t *Tokenizer) fold(position int, depth int, tokens Tokens) (Tokeners, error) {
	if depth >= t.maxDepth {
		return nil, errors.NewErrNestingTooDeep(position, depth)
	}

	var (
		tr     Tokener
		token  Token
		tail   Tokens
		bounds *Bounds
		err    error
	)

	tokensLength := len(tokens)
	buf := Tokeners{}

	for k := 0; k < tokensLength; {
		token = tokens[k]

		switch token {
		case Byte:
			buf = append(buf, token)
		case Boolean:
			buf = append(buf, token)
		case Int16:
			buf = append(buf, token)
		case Int32:
			buf = append(buf, token)
		case Int64:
			buf = append(buf, token)
		case Uint16:
			buf = append(buf, token)
		case Uint32:
			buf = append(buf, token)
		case Uint64:
			buf = append(buf, token)
		case Double:
			buf = append(buf, token)
		case String:
			buf = append(buf, token)
		case ObjectPath:
			buf = append(buf, token)
		case Signature:
			buf = append(buf, token)
		case UnixFD:
			buf = append(buf, token)
		case Array:
			k++
			if k >= tokensLength {
				return nil, NewErrUnexpectedEOF(token)
			}

			tail = tokens[k:]

			if tail[0] != DictStart {
				return nil, NewErrUnexpectedToken(
					position+k,
					tail[0],
				)
			}

			bounds, err = GetBounds(
				position+k,
				tail,
				DictStart,
				DictEnd,
			)
			if err != nil {
				return nil, err
			}

			k += bounds.Starting

			tr, err = t.fold(
				position+k+1,
				depth+1,
				tail[bounds.Starting+1:bounds.Closing],
			)
			if err != nil {
				return nil, err
			}

			buf = append(
				buf,
				&RecursiveToken{
					Tokener: token,
					Head:    DictStart,
					Tail:    DictEnd,
					Body:    tr,
				},
			)

			k += bounds.Closing
		case Variant:
			buf = append(buf, token)
		case Struct:
			return nil, NewErrVirtualToken(position+k, token)
		case StructStart:
			if k+1 >= tokensLength {
				return nil, NewErrUnexpectedEOF(token)
			}

			tail = tokens[k:]

			bounds, err = GetBounds(
				position+k,
				tail,
				StructStart,
				StructEnd,
			)
			if err != nil {
				return nil, err
			}

			k += bounds.Starting

			tr, err = t.fold(
				position+k+1,
				depth+1,
				tail[bounds.Starting+1:bounds.Closing],
			)
			if err != nil {
				return nil, err
			}

			buf = append(
				buf,
				&RecursiveToken{
					Tokener: Struct,
					Head:    StructStart,
					Tail:    StructEnd,
					Body:    tr,
				},
			)

			k += bounds.Closing
		case StructEnd:
			return nil, NewErrUnexpectedToken(position+k, StructEnd)
		case Dict:
			return nil, NewErrVirtualToken(position+k, token)
		case DictStart:
			if k+1 >= tokensLength {
				return nil, NewErrUnexpectedEOF(token)
			}

			tail = tokens[k:]

			bounds, err = GetBounds(
				position+k,
				tail,
				DictStart,
				DictEnd,
			)
			if err != nil {
				return nil, err
			}

			k += bounds.Starting

			tr, err = t.fold(
				position+k+1,
				depth+1,
				tail[bounds.Starting+1:bounds.Closing],
			)
			if err != nil {
				return nil, err
			}

			buf = append(
				buf,
				&RecursiveToken{
					Tokener: Dict,
					Head:    DictStart,
					Tail:    DictEnd,
					Body:    tr,
				},
			)

			k += bounds.Closing
		case DictEnd:
			return nil, NewErrUnexpectedToken(position+k, token)
		default:
			return nil, NewErrUnsupportedToken(position+k, int(token))
		}

		k++
	}
	return buf, nil
}

func NewTokenizer(depth int) *Tokenizer {
	return &Tokenizer{maxDepth: depth}
}

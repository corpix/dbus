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
)

type ErrUnsupportedToken struct {
	token    int
	position int
}

func (e *ErrUnsupportedToken) Error() string {
	return fmt.Sprintf(
		"Unsupported token id: '%d' at position %d",
		e.token,
		e.position,
	)
}

func NewErrUnsupportedToken(position int, token int) error {
	return &ErrUnsupportedToken{token, position}
}

//

type ErrUnexpectedToken struct {
	token    Token
	position int
}

func (e *ErrUnexpectedToken) Error() string {
	return fmt.Sprintf(
		"Unexpected token: '%s' at position %d",
		e.token,
		e.position,
	)
}

func NewErrUnexpectedToken(position int, token Token) error {
	return &ErrUnexpectedToken{token, position}
}

//

type ErrSyntax struct {
	near     string
	position int
}

func (e *ErrSyntax) Error() string {
	return fmt.Sprintf(
		"Syntax error near '%s' at position %d",
		e.near,
		e.position,
	)
}

func NewErrSyntax(position int, near string) error {
	return &ErrSyntax{near, position}
}

//

type ErrBoundNotFound struct {
	starting Token
	closing  Token
}

func (e *ErrBoundNotFound) Error() string {
	return fmt.Sprintf(
		"Bound not found, wanted start token '%s' and close token '%s', found neither of them",
		e.starting.Encode(),
		e.closing.Encode(),
	)
}

func NewErrBoundNotFound(starting, closing Token) error {
	return &ErrBoundNotFound{starting, closing}
}

//

type ErrBoundIncomplete struct {
	starting Token
	closing  Token
	position int
}

func (e *ErrBoundIncomplete) Error() string {
	return fmt.Sprintf(
		"Bound start token '%s' found but close token '%s' is not, bound incomplete at position %d",
		e.starting.Encode(),
		e.closing.Encode(),
		e.position,
	)
}

func NewErrBoundIncomplete(position int, starting, closing Token) error {
	return &ErrBoundIncomplete{starting, closing, position}
}

//

type ErrUnexpectedEOF struct {
	at Token
}

func (e *ErrUnexpectedEOF) Error() string {
	return fmt.Sprintf(
		"Unexpected EOF at position '%s'",
		e.at.Encode(),
	)
}

func NewErrUnexpectedEOF(at Token) error {
	return &ErrUnexpectedEOF{at}
}

//

type ErrVirtualToken struct {
	token    Token
	position int
}

func (e *ErrVirtualToken) Error() string {
	return fmt.Sprintf(
		"Specification error, virtual token '%s' should not appear in signature, but was found at %d, see description of the token at https://dbus.freedesktop.org/doc/dbus-specification.html#idm399",
		e.token.String(),
		e.position,
	)
}

func NewErrVirtualToken(position int, token Token) error {
	return &ErrVirtualToken{token, position}
}

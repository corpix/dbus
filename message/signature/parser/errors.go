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
	"fmt"

	"github.com/corpix/dbus/message/signature/parser/types"
)

type ErrTooMuchTypes struct {
	types    types.Types
	expected int
}

func (e *ErrTooMuchTypes) Error() string {
	return fmt.Sprintf(
		"Too much types: got '%#v' expected %d types",
		e.types,
		e.expected,
	)
}

func NewErrTooMuchTypes(expected int, types types.Types) error {
	return &ErrTooMuchTypes{types, expected}
}

//

type ErrTooFewTypes struct {
	types    types.Types
	expected int
}

func (e *ErrTooFewTypes) Error() string {
	return fmt.Sprintf(
		"Too few types: got '%#v' expected %d types",
		e.types,
		e.expected,
	)
}

func NewErrTooFewTypes(expected int, types types.Types) error {
	return &ErrTooFewTypes{types, expected}
}

//

type ErrTypeExpected struct {
	got interface{}
}

func (e *ErrTypeExpected) Error() string {
	return fmt.Sprintf(
		"Type was expected, got '%#v'",
		e.got,
	)
}

func NewErrTypeExpected(got interface{}) error {
	return &ErrTypeExpected{got}
}

//

type ErrUnsupportedInstance struct {
	t        string
	instance interface{}
}

func (e *ErrUnsupportedInstance) Error() string {
	return fmt.Sprintf(
		"Unsupported instance of '%s' received, got '%T'",
		e.t,
		e.instance,
	)
}

func NewErrUnsupportedInstance(t string, instance interface{}) error {
	return &ErrUnsupportedInstance{t, instance}
}

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
package types

type ComplexType struct {
	Typer
	Body Typer
}

func (t *ComplexType) String() string {
	r := ""

	if t == nil {
		return r
	}

	if t.Typer != nil {
		r += t.Typer.String()
	}

	if t.Body != nil {
		if r != "" {
			r += " "
		}
		r += t.Body.String()
	}

	return r
}

func (t *ComplexType) Kind() Kind {
	return Complex
}

func (t *ComplexType) Types() Types {
	if t == nil || t.Body == nil {
		return Types{}
	}

	return t.Body.Types()
}

func NewComplex(t Typer, body Typer) (*ComplexType, error) {
	if t == nil {
		return nil, ErrTyperIsNil
	}
	return &ComplexType{t, body}, nil
}

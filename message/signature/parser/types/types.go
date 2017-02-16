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

import (
	"fmt"
	"strings"
)

const (
	Invalid Type = iota
	Byte
	Boolean
	Int16
	Int32
	Int64
	Uint16
	Uint32
	Uint64
	Double
	String
	ObjectPath
	Signature
	UnixFD
	Array
	Variant
	Struct
	Dict
)

var (
	TypeNames = map[Type]string{
		Invalid:    "Invalid",
		Byte:       "Byte",
		Boolean:    "Boolean",
		Int16:      "Int16",
		Int32:      "Int32",
		Int64:      "Int64",
		Uint16:     "Uint16",
		Uint32:     "Uint32",
		Uint64:     "Uint64",
		Double:     "Double",
		String:     "String",
		ObjectPath: "ObjectPath",
		Signature:  "Signature",
		UnixFD:     "UnixFD",
		Array:      "Array",
		Variant:    "Variant",
		Struct:     "Struct",
		Dict:       "Dict",
	}
)

type Types []Type

func (ts Types) String() string {
	if ts == nil || len(ts) == 0 {
		return "[ ]"
	}

	buf := make([]string, len(ts))
	for k, v := range ts {
		buf[k] = v.String()
	}
	return fmt.Sprintf(
		"[ %s ]",
		strings.Join(buf, ", "),
	)
}

func (ts Types) Kind() Kind {
	return Complex
}

func (ts Types) Types() Types {
	if ts == nil {
		return Types{}
	}

	return ts
}

func NewTypes(t ...Type) Types {
	return Types(t)
}

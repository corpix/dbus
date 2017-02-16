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

const (
	Invalid Token = iota
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
	StructStart
	StructEnd
	Dict
	DictStart
	DictEnd
)

var (
	EncodedToken = map[rune]Token{
		'y': Byte,
		'b': Boolean,
		'n': Int16,
		'i': Int32,
		'x': Int64,
		'q': Uint16,
		'u': Uint32,
		't': Uint64,
		'd': Double,
		's': String,
		'o': ObjectPath,
		'g': Signature,
		'h': UnixFD,
		'a': Array,
		'v': Variant,
		'r': Struct,
		'(': StructStart,
		')': StructEnd,
		'e': Dict,
		'{': DictStart,
		'}': DictEnd,
	}

	TokenType = map[Token]types.Type{
		Invalid:    types.Invalid,
		Byte:       types.Byte,
		Boolean:    types.Boolean,
		Int16:      types.Int16,
		Int32:      types.Int32,
		Int64:      types.Int64,
		Uint16:     types.Uint16,
		Uint32:     types.Uint32,
		Uint64:     types.Uint64,
		Double:     types.Double,
		String:     types.String,
		ObjectPath: types.ObjectPath,
		Signature:  types.Signature,
		UnixFD:     types.UnixFD,
		Array:      types.Array,
		Variant:    types.Variant,
		Struct:     types.Struct,
		Dict:       types.Dict,
	}

	// TokenVirtual is a workaround for a D-BUS protocol bad design.
	// There are some tokens(`r`, `e`) which are virtual, means
	// they have no Token.Encode() representation in signature and used
	// only for representing some concepts. This map stores this
	// knowledge.
	TokenVirtual = map[Token]bool{
		Struct: true,
		Dict:   true,
	}

	// see init()
	TokenEncoded = map[Token]rune{}
	TypeToken    = map[types.Type]Token{}

	TokenNames = map[Token]string{
		Invalid:     "Invalid",
		Byte:        "Byte",
		Boolean:     "Boolean",
		Int16:       "Int16",
		Int32:       "Int32",
		Int64:       "Int64",
		Uint16:      "Uint16",
		Uint32:      "Uint32",
		Uint64:      "Uint64",
		Double:      "Double",
		String:      "String",
		ObjectPath:  "ObjectPath",
		Signature:   "Signature",
		UnixFD:      "UnixFD",
		Array:       "Array",
		Variant:     "Variant",
		Struct:      "Struct",
		StructStart: "StructStart",
		StructEnd:   "StructEnd",
		Dict:        "Dict",
		DictStart:   "DictStart",
		DictEnd:     "DictEnd",
	}
)

func init() {
	// Precalculate flipped maps of tokens and existing types
	for k, v := range EncodedToken {
		TokenEncoded[v] = k
		if tt, ok := TokenType[v]; ok {
			TypeToken[tt] = v
		}
	}
}

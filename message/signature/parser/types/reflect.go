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
	"reflect"
)

var (
	ReflectType = map[Type]reflect.Kind{
		Byte:    reflect.Uint64,
		Boolean: reflect.Bool,
		Int16:   reflect.Int16,
		Int32:   reflect.Int32,
		Int64:   reflect.Int64,
		Uint16:  reflect.Uint16,
		Uint32:  reflect.Uint32,
		Uint64:  reflect.Uint64,
		Double:  reflect.Complex128,
		String:  reflect.String,
		Array:   reflect.Slice,
		Struct:  reflect.Struct,
		Dict:    reflect.Map,
	}
)

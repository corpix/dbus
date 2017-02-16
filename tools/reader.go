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
package tools

import (
	"bytes"
	"io"
)

func ReadBytes(r io.Reader, terminator []byte) ([]byte, error) {
	var (
		b             []byte
		n             int
		delimiter     int
		lastIteration bool
		err           error
	)

	buf := []byte{}

	for {
		b = make([]byte, bytes.MinRead)
		n, err = r.Read(b)
		if err != nil {
			if err != io.EOF {
				lastIteration = true
			} else {
				return nil, err
			}
		}
		if n > 0 {
			delimiter = bytes.Index(b, terminator)
			if delimiter >= 0 {
				lastIteration = true
			} else {
				delimiter = len(b)
			}
			buf = append(buf, b[:delimiter]...)
		}
		if lastIteration {
			break
		}
	}

	// `err` could contain io.EOF
	// It should be returned to indicate
	// the end of a stream.
	return buf, err
}

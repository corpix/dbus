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
	"github.com/corpix/dbus/message/signature/parser/tokenizer"
	"github.com/corpix/dbus/message/signature/parser/types"
)

func singleType(ts types.Types) (*types.Type, error) {
	if len(ts) == 0 {
		return nil, NewErrTooFewTypes(
			1,
			ts,
		)
	}
	if len(ts) > 1 {
		return nil, NewErrTooMuchTypes(
			1,
			ts,
		)
	}

	t := ts[0]
	return &t, nil
}

func RecursiveTokenToTyper(tokener *tokenizer.RecursiveToken) (types.Typer, error) {
	if tokener.Tokener == nil {
		return nil, NewErrTypeExpected(tokener.Tokener)
	}

	t, err := singleType(tokener.Tokener.Types())
	if err != nil {
		return nil, err
	}

	content, err := TokenerToTyper(tokener.Body)
	if err != nil {
		return nil, err
	}

	return types.NewComplex(t, content)
}

func TokenersToTyper(tokeners tokenizer.Tokeners) (types.Typer, error) {
	buf := make(types.Typers, len(tokeners))
	var (
		typer types.Typer
		err   error
	)
	for k, v := range tokeners {
		typer, err = TokenerToTyper(v)
		if err != nil {
			return nil, err
		}

		buf[k] = typer
	}
	return buf, nil
}

func TokensToTyper(tokens tokenizer.Tokens) (types.Typer, error) {
	buf := make(types.Types, len(tokens))
	var (
		t   *types.Type
		err error
	)
	for k, v := range tokens {
		t, err = singleType(v.Types())
		if err != nil {
			return nil, err
		}
		buf[k] = *t
	}

	return buf, nil
}

func TokenToTyper(token tokenizer.Token) (types.Typer, error) {
	t, err := singleType(token.Types())
	if err != nil {
		return nil, err
	}

	return types.NewSimple(*t), nil
}

func TokenerToTyper(tokener tokenizer.Tokener) (types.Typer, error) {
	if tokener == nil {
		return nil, NewErrTypeExpected(tokener)
	}

	switch tokener.(type) {
	case *tokenizer.RecursiveToken:
		return RecursiveTokenToTyper(tokener.(*tokenizer.RecursiveToken))
	case tokenizer.Tokeners:
		return TokenersToTyper(tokener.(tokenizer.Tokeners))
	case tokenizer.Tokens:
		return TokensToTyper(tokener.(tokenizer.Tokens))
	case tokenizer.Token:
		return TokenToTyper(tokener.(tokenizer.Token))
	default:
		return nil, NewErrUnsupportedInstance(
			"tokenizer.Tokener",
			tokener,
		)
	}
}

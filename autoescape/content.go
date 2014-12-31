// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package autoescape

import "github.com/robfig/soy/data"

// Strings of content from a trusted source.
// These types are data.Values.
type (
	// CSS encapsulates known safe content that matches any of:
	//   1. The CSS3 stylesheet production, such as `p { color: purple }`.
	//   2. The CSS3 rule production, such as `a[href=~"https:"].foo#bar`.
	//   3. CSS3 declaration productions, such as `color: red; margin: 2px`.
	//   4. The CSS3 value production, such as `rgba(0, 0, 255, 127)`.
	// See http://www.w3.org/TR/css3-syntax/#parsing and
	// https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style
	CSS string

	// HTML encapsulates a known safe HTML document fragment.
	// It should not be used for HTML from a third-party, or HTML with
	// unclosed tags or comments. The outputs of a sound HTML sanitizer
	// and a template escaped by this package are fine for use with HTML.
	HTML string

	// HTMLAttr encapsulates an HTML attribute from a trusted source,
	// for example, ` dir="ltr"`.
	HTMLAttr string

	// JS encapsulates a known safe EcmaScript5 Expression, for example,
	// `(x + y * z())`.
	// Template authors are responsible for ensuring that typed expressions
	// do not break the intended precedence and that there is no
	// statement/expression ambiguity as when passing an expression like
	// "{ foo: bar() }\n['foo']()", which is both a valid Expression and a
	// valid Program with a very different meaning.
	JS string

	// JSStr encapsulates a sequence of characters meant to be embedded
	// between quotes in a JavaScript expression.
	// The string must match a series of StringCharacters:
	//   StringCharacter :: SourceCharacter but not `\` or LineTerminator
	//                    | EscapeSequence
	// Note that LineContinuations are not allowed.
	// JSStr("foo\\nbar") is fine, but JSStr("foo\\\nbar") is not.
	JSStr string

	// URL encapsulates a known safe URL or URL substring (see RFC 3986).
	// A URL like `javascript:checkThatFormNotEditedBeforeLeavingPage()`
	// from a trusted source should go in the page, but by default dynamic
	// `javascript:` URLs are filtered out since they are a frequently
	// exploited injection vector.
	URL string
)

func (v HTML) Truthy() bool     { return v != "" }
func (v HTMLAttr) Truthy() bool { return v != "" }
func (v JS) Truthy() bool       { return v != "" }
func (v JSStr) Truthy() bool    { return v != "" }
func (v URL) Truthy() bool      { return v != "" }
func (v CSS) Truthy() bool      { return v != "" }

func (v HTML) String() string     { return string(v) }
func (v HTMLAttr) String() string { return string(v) }
func (v JS) String() string       { return string(v) }
func (v JSStr) String() string    { return string(v) }
func (v URL) String() string      { return string(v) }
func (v CSS) String() string      { return string(v) }

func (v HTML) Equals(other data.Value) bool {
	o, ok := other.(HTML)
	return ok && string(v) == string(o)
}

func (v HTMLAttr) Equals(other data.Value) bool {
	o, ok := other.(HTMLAttr)
	return ok && string(v) == string(o)
}

func (v JS) Equals(other data.Value) bool {
	o, ok := other.(JS)
	return ok && string(v) == string(o)
}

func (v JSStr) Equals(other data.Value) bool {
	o, ok := other.(JSStr)
	return ok && string(v) == string(o)
}

func (v URL) Equals(other data.Value) bool {
	o, ok := other.(URL)
	return ok && string(v) == string(o)
}

func (v CSS) Equals(other data.Value) bool {
	o, ok := other.(CSS)
	return ok && string(v) == string(o)
}

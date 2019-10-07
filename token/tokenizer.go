package token

import (
	"bytes"
	"fmt"
	"io"
)

type Token struct {
	Type     TokenType
	Alias    []TokenType
	Content  string
	Position Position
}

// A source position is represented by a Position value.
// A position is valid if Line > 0.
type Position struct {
	Filename string // filename, if any
	Offset   int    // byte offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (character count per line)
}

// IsValid reports whether the position is valid.
func (pos *Position) IsValid() bool { return pos.Line > 0 }

func (pos Position) String() string {
	s := pos.Filename
	if s != "" {
		s += ":"
	}
	if pos.IsValid() {
		s = fmt.Sprintf("%s%d:%d", s, pos.Line, pos.Column)
	}
	return s
}

const bufLen = 1024 // at least utf8.UTFMax

type Tokenizer struct {
	tokenset *TokenSet

	// Input
	src io.Reader

	// Source buffer
	srcBuf [bufLen + 1]byte // +1 for sentinel for common case of s.next()
	srcPos int              // reading position (srcBuf index)
	srcEnd int              // source end (srcBuf index)

	// Source position
	srcBufOffset int // byte offset of srcBuf[0] in source
	line         int // line count
	column       int // character count
	lastLineLen  int // length of last line in characters (for correct column reporting)
	lastCharLen  int // length of last character in bytes

	// Token text buffer
	// Typically, token text is stored completely in srcBuf, but in general
	// the token text's head may be buffered in tokBuf while the token text's
	// tail is stored in srcBuf.
	tokBuf bytes.Buffer // token text head that is not in srcBuf anymore
	tokPos int          // token text tail position (srcBuf index); valid if >= 0
	tokEnd int          // token text tail end (srcBuf index)

	// One character look-ahead
	ch rune // character before current srcPos

	// Start position of most recently scanned token; set by Scan.
	// Calling Init or Next invalidates the position (Line == 0).
	// The Filename field is always left untouched by the Scanner.
	// If an error is reported (via Error) and Position is invalid,
	// the scanner is not inside a token. Call Pos to obtain an error
	// position in that case, or to obtain the position immediately
	// after the most recently scanned token.
	Position
}

func (tok *Tokenizer) Init(filename string, reader io.Reader, tokenset *TokenSet) {

}

func (tok *Tokenizer) InitByText(text []byte, tokenset *TokenSet) {

}

func (tok *Tokenizer) Next() Token {
	return Token{}
}

func (tok *Tokenizer) Pos() Position {
	return tok.Position
}

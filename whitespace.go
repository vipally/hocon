package hocon

import (
	"bytes"
)

/*
SPACE (\u0020)
NO-BREAK SPACE (\u00A0)
OGHAM SPACE MARK (\u1680)
EN QUAD (\u2000)
EM QUAD (\u2001)
EN SPACE (\u2002)
EM SPACE (\u2003)
THREE-PER-EM SPACE (\u2004)
FOUR-PER-EM SPACE (\u2005)
SIX-PER-EM SPACE (\u2006)
FIGURE SPACE (\u2007)
PUNCTUATION SPACE (\u2008)
THIN SPACE (\u2009)
HAIR SPACE (\u200A)
NARROW NO-BREAK SPACE (\u202F)
MEDIUM MATHEMATICAL SPACE (\u205F)
and IDEOGRAPHIC SPACE (\u3000)
Byte Order Mark (\uFEFF)
*/

var whiteSpaces = [][]byte{
	[]byte(" "),
	[]byte("\t"),
	[]byte("\n"),
	[]byte("\u000B"),
	[]byte("\u000C"),
	[]byte("\u000D"),
	[]byte("\u00A0"),
	[]byte("\u1680"),
	[]byte("\u2000"),
	[]byte("\u2001"),
	[]byte("\u2002"),
	[]byte("\u2003"),
	[]byte("\u2004"),
	[]byte("\u2005"),
	[]byte("\u2006"),
	[]byte("\u2007"),
	[]byte("\u2008"),
	[]byte("\u2009"),
	[]byte("\u200A"),
	[]byte("\u202F"),
	[]byte("\u205F"),
	[]byte("\u2060"),
	[]byte("\u3000"),
	[]byte("\uFEFF"),
}

func isWhitespace(b []byte) (int, bool) {
	if buffsize := len(b); buffsize > 0 {
		switch b[0] {
		case 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x20, 0xc2, 0xe1, 0xe2, 0xe3, 0xef:
			for _, v := range whiteSpaces {
				if size := len(v); buffsize >= size {
					if bytes.Compare(v, b[:size]) == 0 {
						return size, true
					}
				}
			}
		}
	}
	return 0, false
}

func skipWhitespace(b []byte) []byte {
	for {
		if n, ok := isWhitespace(b); ok {
			b = b[n:]
		} else {
			break
		}
	}
	return b
}

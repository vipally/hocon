//custom token define
package token

import (
	"errors"
	"fmt"
	"io"
)

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

type FuncMatchToken func(ch rune, index int) bool

type TokenDefineElem struct {
	token        TokenType
	name         string
	schemas      []string
	firstRunes   map[rune]*[][]byte
	fnMatch      FuncMatchToken
	expectEnd    TokenType
	contentToken TokenType
	expectNext   []TokenType
	probAlias    []TokenType
}

func NewTokenset(name string) *Tokenset {
	p := &Tokenset{}
	p.Init(name)
	return p
}

type Tokenset struct {
	locked     bool                        //lock for redister
	name       string                      //name of TokenSet
	tokens     []*TokenDefineElem          //token defines
	fastToken  map[rune][]*TokenDefineElem //fast entry of token by rune
	otherToken []*TokenDefineElem          //tokens matches no schema,eg: ident, number
}

func (ts *Tokenset) Init(name string) error {
	if ts.locked {
		return errors.New("Init on locked Tokenset")
	}
	if ts.name != "" {
		return errors.New("Init on non-empty Tokenset")
	}
	ts.name = name
	ts.tokens = make([]*TokenDefineElem, int(MinTokenTypeRegister), 8)
	ts.fastToken = make(map[rune][]*TokenDefineElem)
	ts.otherToken = make([]*TokenDefineElem, 0, 2)
	//ts.Register(TokenEOF, "EOF", "")
	//ts.Register(TokenNewline, "NewLine", "\n")
	return nil
}

func (ts *Tokenset) getTokenObj(token TokenType, alloc bool) (*TokenDefineElem, error) {
	var obj *TokenDefineElem
	idx := int(token - MinTokenTypeRegister)
	switch {
	case !token.IsRegisterToken():
		return nil, fmt.Errorf("invalid token %d", token)
	default:
		return nil, fmt.Errorf("tokenset grow up fail:%d", token)
	case idx < len(ts.tokens):
		obj = ts.tokens[idx]
	case idx == len(ts.tokens) && alloc:
		ts.tokens = append(ts.tokens, &TokenDefineElem{})
	}
	return obj, nil
}

func (ts *Tokenset) RegisterFunc(token TokenType, name string, fnMatch FuncMatchToken) error {
	if ts.locked {
		return errors.New("Register on locked Tokenset")
	}
	return nil
}

func (ts *Tokenset) Register(token TokenType, name string, schema string) error {
	if ts.locked {
		return errors.New("Register on locked Tokenset")
	}
	return nil
}

func (ts *Tokenset) RegExpectedPair(startToken, endToken, contentToken TokenType) error {
	if ts.locked {
		return errors.New("Register on locked Tokenset")
	}
	return nil
}

func (ts *Tokenset) RegExpectedNext(startToken TokenType, nextTokens ...TokenType) error {
	if ts.locked {
		return errors.New("Register on locked Tokenset")
	}
	return nil
}

func (ts *Tokenset) RegProbAlias(aliasTokens ...TokenType) error {
	if ts.locked {
		return errors.New("Register on locked Tokenset")
	}
	return nil
}

func (ts *Tokenset) Lock() {
	ts.locked = true
}

func (ts *Tokenset) Name() string {
	return ts.name
}

func (ts *Tokenset) TokenName(token TokenType) string {
	if tok, err := ts.getTokenObj(token, false); err == nil {
		return tok.name
	}
	return ""
}

func (ts *Tokenset) NewParser(filename string, reader io.Reader) *Scanner {
	return nil
}

func (ts *Tokenset) NewParserByText(text []byte) *Scanner {
	return nil
}

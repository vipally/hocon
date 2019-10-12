//sustom token define
package token

import (
	"io"
)

const (
	minTokenType           = 1
	MaxTokenType           = 65535
	invalidToken           = -1
	TokenNone    TokenType = 0
)

const (
	TokenEOF TokenType = iota + minTokenType
	TokenNewline
	TokenWhiteSpace

	MinTokenType
)

func (token TokenType) IsValid() bool {
	return token >= minTokenType && token <= MaxTokenType
}

func (token TokenType) IsZero() bool {
	return token == TokenNone
}

func (token TokenType) Int() int {
	if token >= 0 && token <= MaxTokenType {
		return int(token)
	}
	return invalidToken
}

// TokenType represents the enum of tokens
type TokenType int

type Peaker interface {
	Peak() rune
}

type FuncMatchToken func(peaker Peaker) []byte

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

type TokenSet struct {
	locked     bool               //lock for redister
	name       string             //name of TokenSet
	tokens     []*TokenDefineElem //token defines
	fastToken  map[rune][]*TokenDefineElem
	otherToken []*TokenDefineElem
}

func (ts *TokenSet) RegisterFunc(token TokenType, name string, fnMatch FuncMatchToken) error {
	return nil
}

func (ts *TokenSet) Register(token TokenType, name string, schemas ...string) error {
	return nil
}

func (ts *TokenSet) RegExpectedPair(startToken, endToken, contentToken TokenType) error {
	return nil
}

func (ts *TokenSet) RegExpectedNext(startToken TokenType, nextTokens ...TokenType) error {
	return nil
}

func (ts *TokenSet) RegProbAlias(aliasTokens ...TokenType) error {
	return nil
}

func (ts *TokenSet) Lock() {
	ts.locked = true
}

func (ts *TokenSet) Name() string {
	return ts.name
}

func (ts *TokenSet) TokenName(token TokenType) string {
	return ""
}

func (ts *TokenSet) NewTokenizer(filename string, reader io.Reader) *Scanner {
	return nil
}

func (ts *TokenSet) NewTokenizerByText(text []byte) *Scanner {
	return nil
}

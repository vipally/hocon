package hocon

import (
	"fmt"
)

type tokenType int

const (
	tokenTypeNone         tokenType = iota //
	tokenTypeComment                       // "//" or "#"
	tokenTypeKey                           //
	tokenTypeLiteralValue                  //
	tokenTypeAssign                        // "=" or ":"
	tokenTypePlusAssign                    // +=
	tokenTypeObjectStart                   // {
	tokenTypeObjectEnd                     // }
	tokenTypeDot                           // .
	tokenTypeNewline                       // \n
	tokenTypeEOF                           //
	tokenTypeArrayStart                    // [
	tokenTypeArrayEnd                      // ]
	tokenTypeComma                         // ,
	tokenTypeSubstitute                    // $
	tokenTypeInclude                       // include

	_tokenTypeNum
)

var tokentypeName = [_tokenTypeNum]string{
	"tokenTypeNone",
	"tokenTypeComment",
	"tokenTypeKey",
	"tokenTypeLiteralValue",
	"tokenTypeAssign",
	"tokenTypePlusAssign",
	"tokenTypeObjectStart",
	"tokenTypeObjectEnd",
	"tokenTypeDot",
	"tokenTypeNewline",
	"tokenTypeEOF",
	"tokenTypeArrayStart",
	"tokenTypeArrayEnd",
	"tokenTypeComma",
	"tokenTypeSubstitute",
	"tokenTypeInclude",
}

func (token tokenType) Valid() bool {
	return token >= 0 && token < _tokenTypeNum
}

func (token tokenType) String() string {
	if token.Valid() {
		return tokentypeName[token]
	}
	return fmt.Sprintf("undefinedToken%d", token)
}

type token struct {
	tokenType tokenType
	value     []byte
	optional  bool
}

func newToken(v interface{}) *token {
	switch value := v.(type) {
	case []byte:
		{
			return &token{tokenType: tokenTypeLiteralValue, value: value}
		}
	case tokenType:
		{
			return &token{tokenType: value}
		}
	}

	return nil
}

func newKeyToken(key []byte) *token {
	return &token{tokenType: tokenTypeKey, value: key}
}

func newSubstitutionToken(path []byte, optional bool) *token {
	return &token{tokenType: tokenTypeSubstitute, value: path, optional: optional}
}

func newLiteralValueToken(value []byte) *token {
	return &token{tokenType: tokenTypeLiteralValue, value: value}
}

func newIncludeToken(path []byte) *token {
	return &token{tokenType: tokenTypeInclude, value: path}
}

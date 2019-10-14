package token

const (
	minTokenType = 1
	MaxTokenType = 65535
	invalidToken = -1
)

const (
	TokenNone            TokenType = 0                   //
	TokenEOF             TokenType = iota + minTokenType //end of file
	TokenNewline                                         //new line
	MinTokenTypeRegister                                 //The first TokenType for register outside
)

func (token TokenType) IsValid() bool {
	return token >= minTokenType && token <= MaxTokenType
}

func (token TokenType) IsZero() bool {
	return token == TokenNone
}

func (token TokenType) IsTokenEOF() bool {
	return token == TokenEOF
}

func (token TokenType) IsTokenNewLine() bool {
	return token == TokenNewline
}

func (token TokenType) IsRegisterToken() bool {
	return token.IsValid() && token >= MinTokenTypeRegister
}

func (token TokenType) ToInt() int {
	if token >= 0 && token <= MaxTokenType {
		return int(token)
	}
	return invalidToken
}

// TokenType represents the enum of tokens
type TokenType int

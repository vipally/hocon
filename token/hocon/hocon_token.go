package hocon

import (
	"github.com/vipally/hocon/token"
)

var tok = token.NewTokenset("HOCON")

func Token() *token.Tokenset {
	return tok
}

func init() {
	//Token.Register()
}

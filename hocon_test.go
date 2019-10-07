package hocon

import (
	"fmt"
	"testing"
	"unicode"
)

func _TestWhiteSpace(t *testing.T) {
	for i, v := range whiteSpaces {
		fmt.Printf("%02d [%s] %02x %x\n", i, string(v), v[0], v)
	}
}

func TestUtf8(t *testing.T) {
	//fmt.Printf("%x\n", unicode.MaxRune)
	for i := rune(0); i <= unicode.MaxRune*0+0x7ff; i++ {
		s := string(i)
		fmt.Printf("\\u%04x %x %s\n", i, []byte(s), s)
	}
}

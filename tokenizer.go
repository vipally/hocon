package hocon

var (
	hoconNotInUnquotedKey  = []byte("$\"{}[]:=+,#`^?!@*&\\.")
	hoconNotInUnquotedText = []byte("$\"{}[]:=+,#`^?!@*&\\")
)

type tokenizer struct {
	text  []byte
	pos   int
	line  int
	col   int
	stack *stack
}

func newTokenizer(text []byte) *tokenizer {
	return &tokenizer{
		stack: newStack(),
		text:  text,
	}
}

func (p *tokenizer) Push() {
	p.stack.Push(p.pos, p.line, p.col)
}

func (p *tokenizer) Pop() error {
	pos, line, col, err := p.stack.Pop()
	if err != nil {
		return err
	}
	p.pos, p.line, p.col = pos, line, col
	return nil
}

func (p *tokenizer) EOF() bool {
	return p.pos >= len(p.text)
}

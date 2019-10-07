package hocon

import (
	"errors"
)

type stackElem struct {
	pos  int // pos of data
	line int // line number
	col  int // col in line
}

type stack struct {
	data []stackElem
}

func newStack() *stack {
	s := &stack{
		data: make([]stackElem, 0, 10),
	}
	return s
}

func (st *stack) Push(pos, line, col int) {
	elem := stackElem{
		pos:  pos,
		line: line,
		col:  col,
	}
	st.data = append(st.data, elem)
}

func (p *stack) Pop() (int, int, int, error) {
	size := len(p.data)
	if size == 0 {
		return 0, 0, 0, errors.New("empty stack")
	}

	d := p.data[size-1]
	p.data = p.data[:size-1]
	return d.pos, d.line, d.col, nil
}

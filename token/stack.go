package token

import (
	"errors"
)

type TokenPosition struct {
	TokenType TokenType
	Content   string
	Pos       Position
}

type stack struct {
	data []TokenPosition
}

func newStack() *stack {
	s := &stack{
		data: make([]TokenPosition, 0, 16),
	}
	return s
}

func (st *stack) Push(pos TokenPosition) {
	st.data = append(st.data, pos)
}

func (st *stack) Pop() (TokenPosition, error) {
	size := len(st.data)
	if size == 0 {
		return TokenPosition{}, errors.New("empty stack")
	}

	d := st.data[size-1]
	st.data = st.data[:size-1]
	return d, nil
}

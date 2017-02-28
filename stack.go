// Package gostack provides a general purpose stack type
package gostack

// Stack represents a stack of objects
type Stack struct {
	stack []interface{}
}

// Count returns the number of items on the stack
func (s *Stack) Count() int {

	return len(s.stack)
}

// Peep returns the top item on the stack without removing it
func (s *Stack) Peep() (interface{}, bool) {

	l := len(s.stack)

	if l == 0 {
		return nil, false
	}

	return s.stack[l-1], true
}

// Pop removes the top item from the stack and returns it
func (s *Stack) Pop() (interface{}, bool) {

	l := len(s.stack)

	if l == 0 {
		return nil, false
	}

	v := s.stack[l-1]
	s.stack = s.stack[:l-1]

	return v, true
}

// Push places a new item on the top of the stack
func (s *Stack) Push(v interface{}) {

	s.stack = append(s.stack, v)
}

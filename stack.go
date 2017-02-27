package gostack

type Stack struct {
	stack []interface{}
}

func (s *Stack) Count() int {

	return len(s.stack)
}

func (s *Stack) Peep() (interface{}, bool) {

	l := len(s.stack)

	if l == 0 {
		return nil, false
	}

	return s.stack[l-1], true
}

func (s *Stack) Pop() (interface{}, bool) {

	l := len(s.stack)

	if l == 0 {
		return nil, false
	}

	v := s.stack[l-1]
	s.stack = s.stack[:l-1]

	return v, true
}

func (s *Stack) Push(v interface{}) {

	s.stack = append(s.stack, v)
}

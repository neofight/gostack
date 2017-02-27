package gostack

import (
	"testing"
)

func TestEmptyStackCount(t *testing.T) {

	var s Stack

	if s.Count() != 0 {
		t.Error("Expected empty stack to have a count of 0")
	}
}

func TestEmptyStackPeep(t *testing.T) {

	var s Stack

	v, ok := s.Peep()

	if v != nil {
		t.Error("Expected peep of empty stack to return an empty interface{}")
	}

	if ok != false {
		t.Error("Expected peep of empty stack to indicate that the operation had failed")
	}
}

func TestEmptyStackPop(t *testing.T) {

	var s Stack

	v, ok := s.Pop()

	if v != nil {
		t.Error("Expected pop of empty stack to return an empty interface{}")
	}

	if ok != false {
		t.Error("Expected pop of empty stack to indicate that the operation had failed")
	}
}

func TestStackPushAndCount(t *testing.T) {

	var s Stack

	for i := 0; i < 10; i++ {
		s.Push(i)

		if s.Count() != i+1 {
			t.Error("Count does not match number of items pushed onto the stack")
		}
	}
}

func TestStackPushAndPeep(t *testing.T) {
	var s Stack

	for i := 0; i < 10; i++ {
		s.Push(i)

		v, ok := s.Peep()

		if v != i {
			t.Error("Peep does not match the last item pushed onto the stack")
		}

		if ok != true {
			t.Error("Expected peep of non empty stack to indicate that the operation had succeeded")
		}
	}
}

func TestStackPushAndPop(t *testing.T) {
	var s Stack

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := 9; i >= 0; i-- {

		v, ok := s.Pop()

		if v != i {
			t.Errorf("Expected to pop value %v but saw value %v", i, v)
		}

		if ok != true {
			t.Error("Expected pop of non empty stack to indicate that the operation had succeeded")
		}
	}
}

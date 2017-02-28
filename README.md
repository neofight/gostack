# gostack
The package gostack provides a general purpose stack type.

The usual Count(), Peep(), Pop() and Push() operations are supported. The empty interface (interface{}) is used to allow any type be placed on and retrieved from the stack.

If the stack is not empty Peep() and Pop() both return a value from the stack and a boolean value indicating success. In the event of an empty stack a nil value and a boolean value indicating failure are returned.

The following code taken from one of the tests gives an example of how to use the stack object:

```golang
	var s Stack

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := 9; i >= 0; i-- {

		v, ok := s.Pop()

		if v != i {
			t.Errorf("Expected to pop value %v but saw value %v", i, v)
		}

		if !ok {
			t.Error("Expected pop of non empty stack to indicate that the operation had succeeded")
		}
	}
```

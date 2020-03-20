package stack

import (
	"errors"
	"reflect"
)

// Stack implementation based on slise of
// interface type for using with
// various data type
type Stack []interface{}

// IsEmpty check
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push new data to top of the Stack
func (s *Stack) Push(e interface{}) error {
	if !(*s).IsEmpty() && (reflect.TypeOf((*s)[0]) != reflect.TypeOf(e)) {
		ts := reflect.TypeOf((*s)[0])
		te := reflect.TypeOf(e)
		return errors.New("invalid operation: mismatched types " + ts.Name() + " and " + te.Name())
	}
	*s = append(*s, e)
	return nil
}

// Pop remove and return top element
// of Stack. Return (nil,false) if Stack is empty.
func (s *Stack) Pop() (interface{}, error) {
	if (*s).IsEmpty() {
		return nil, errors.New("empty Stack")
	}
	end := len(*s) - 1
	item := (*s)[end]
	*s = (*s)[:end]
	return item, nil
}

// Rotate the elements of the stack so
// that the first becomes the last
// and the last becomes the first.
func (s *Stack) Rotate() {
	start := 0
	end := len(*s) - 1
	for start = 0; start < len(*s)/2; {
		(*s)[start], (*s)[end] = (*s)[end], (*s)[start]
		start++
		end--
	}
}

// Shift all elements of stack
// by 1, the first element becomes the last one.
func (s *Stack) Shift() {
	first, err := (*s).Pop()
	if err != nil {
		return
	}
	var temp Stack
	temp = append(temp, first)
	(*s) = append(temp, (*s)...)
}

// Unshift all elements of stack
// by 1, the last element becomes the first one.
func (s *Stack) Unshift() {
	if s.IsEmpty() {
		return
	}
	temp := (*s)[0]
	(*s) = (*s)[1:]
	(*s) = append((*s), temp)
}

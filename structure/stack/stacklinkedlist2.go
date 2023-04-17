// Stack Linked-List implementation with a standard library (Container/List).
// A stack is a linear data structure that follows a certain order of operations.
// The order can be LIFO (Last In First Out) or FILO (First In Last Out).
// Reference:
// 	 Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	 Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)

package stack

import (
	"container/list"
	"errors"
)

// SList is struct that point to stack with container/list.List library
type SList struct {
	stack *list.List
}

// Push add a value into stack
func (sl *SList) Push(val any) {
	sl.stack.PushFront(val)
}

// Peak is return last value that insert into stack
func (sl *SList) Peak() (any, error) {
	if sl.Empty() {
		return "", errors.New("stack list is empty")
	}
	element := sl.stack.Front()
	return element.Value, nil
}

// Pop returns the last value that was inserted on the stack and
// also removes it from the stack.
func (sl *SList) Pop() (any, error) {
	if sl.Empty() {
		return "", errors.New("stack list is empty")
	}

	// get last element that insert into stack
	element := sl.stack.Front()
	// remove element in stack
	sl.stack.Remove(element)
	// return element value
	return element.Value, nil
}

// Length return length of stack.
func (sl *SList) Length() int {
	return sl.stack.Len()
}

// Empty checks if the stack has a value or not,
// if 0, then stack is empty, otherwise it is not empty.
func (sl *SList) Empty() bool {
	return sl.stack.Len() == 0
}

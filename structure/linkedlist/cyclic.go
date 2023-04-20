package linkedlist

import "fmt"

// Cyclic is the structure that cycles the linked list in this implementation.
type Cyclic[T any] struct {
	Size int
	Head *Node[T]
}

func NewCyclic[T any]() *Cyclic[T] {
	return &Cyclic[T]{}
}

// Inserting the first node is a special case.
// It will point to itself.
// In other cases the node will be added to the end of the list.
// The end of the list is the Prev field of the current item.
// The complexity is O(1).
func (cl *Cyclic[T]) Add(val T) {
	n := NewNode(val)
	cl.Size++

	if cl.Head == nil {
		n.Prev = n
		n.Next = n
		cl.Head = n
	} else {
		n.Prev = cl.Head.Prev
		n.Next = cl.Head
		cl.Head.Prev.Next = n
		cl.Head.Prev = n
	}
}

// Rotate rotates the list by P places.
// This method is interesting for optimization.
// The first optimization is to reduce the value of P so that it is between 0 and N-1.
// To do this, we have to use the division by modulo operation.
// But be careful if P is less than 0. If it is, make it positive.
// You can do this without breaking the sense of the number by adding a multiple of N to it.
// Now you can reduce P modulo N to rotate the list by the minimum number of places.
// We use the fact that moving forward in a circle by
// P places is equal to moving backward by N - P places.
// So if P > N / 2, we can rotate the list backwards by N - P places.
// Complexity O(n).
func (cl *Cyclic[T]) Rotate(places int) {
	if cl.Size > 0 {
		if places < 0 {
			multiple := cl.Size - 1 - places/cl.Size
			places += multiple * cl.Size
		}
		places %= cl.Size

		if places > cl.Size/2 {
			places = cl.Size - places
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Prev
			}
		} else if places == 0 {
			return
		} else {
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Next
			}

		}
	}
}

// Delete deletes the current item.
func (cl *Cyclic[T]) Delete() bool {
	var deleted bool
	var prevItem, thisItem, nextItem *Node[T]

	if cl.Size == 0 {
		return deleted
	}

	deleted = true
	thisItem = cl.Head
	nextItem = thisItem.Next
	prevItem = thisItem.Prev

	if cl.Size == 1 {
		cl.Head = nil
	} else {
		cl.Head = nextItem
		nextItem.Prev = prevItem
		prevItem.Next = nextItem
	}
	cl.Size--

	return deleted
}

// Destroy destroys all items in the list.
func (cl *Cyclic[T]) Destroy() {
	for cl.Delete() {
		continue
	}
}

// Walk shows list body.
func (cl *Cyclic[T]) Walk() *Node[T] {
	var start *Node[T]
	start = cl.Head

	for i := 0; i < cl.Size; i++ {
		fmt.Printf("%v \n", start.Val)
		start = start.Next
	}
	return start
}

// This is a struct-based solution for Josephus problem.
// Reference: https://en.wikipedia.org/wiki/Josephus_problem
func JosephusProblem(cl *Cyclic[int], k int) int {
	for cl.Size > 1 {
		cl.Rotate(k)
		cl.Delete()
		cl.Rotate(-1)
	}

	retval := cl.Head.Val
	cl.Destroy()

	return retval
}

// Queue Linked-List with standard library (Container/List)
// Queue is a linear structure which follows a particular order in which the operations are performed.
// The order is First In First Out (FIFO).
// Reference: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)

package queue

import (
	"container/list" // used as linked-list
	"errors"
)

// LQueue will be store the value into the list.
type LQueue struct {
	queue *list.List
}

// Enqueue will be added new value
func (lq *LQueue) Enqueue(value any) {
	lq.queue.PushBack(value)
}

// Dequeue will be removed the first value that input
// (First In First Out - FIFO).
func (lq *LQueue) Dequeue() error {

	if !lq.Empty() {
		element := lq.queue.Front()
		lq.queue.Remove(element)

		return nil
	}

	return errors.New("dequeue is empty we got an error")
}

// Front it will return the front value
func (lq *LQueue) Front() (any, error) {
	if !lq.Empty() {
		val := lq.queue.Front().Value
		return val, nil
	}

	return "", errors.New("error queue is empty")
}

// Back it will return the back value
func (lq *LQueue) Back() (any, error) {
	if !lq.Empty() {
		val := lq.queue.Back().Value
		return val, nil
	}

	return "", errors.New("error queue is empty")
}

// Len it will return the length of list
func (lq *LQueue) Len() int {
	return lq.queue.Len()
}

// Empty is check our list is empty or not
func (lq *LQueue) Empty() bool {
	return lq.queue.Len() == 0
}

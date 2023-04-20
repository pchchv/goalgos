package queue

import (
	"testing"
)

func TestQueueLinkedList(t *testing.T) {
	t.Run("Test Queue Linked List", func(t *testing.T) {

		t.Run("Test EnQueue", func(t *testing.T) {
			var newQueue Queue
			newQueue.enqueue(2)
			newQueue.enqueue(3)
			newQueue.enqueue(4)
			newQueue.enqueue(45)

			if newQueue.frontQueue() != 2 && newQueue.backQueue() != 45 {
				t.Errorf("Test EnQueue is wrong the result must be %v and %v but got %v and %v", 2, 45, newQueue.frontQueue(), newQueue.backQueue())
			}

		})

		t.Run("Test DeQueue", func(t *testing.T) {
			var newQueue Queue
			newQueue.enqueue(2)
			newQueue.enqueue(3)
			newQueue.enqueue(4)

			newQueue.dequeue()
			if newQueue.dequeue() != 3 {
				t.Errorf("Test DeQueue is wrong the result must be %v but got %v", 3, newQueue.dequeue())
			}
		})

		t.Run("Test Queue isEmpty", func(t *testing.T) {
			var newQueue Queue
			if newQueue.isEmpty() != true {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", true, newQueue.isEmpty())
			}

			newQueue.enqueue(3)
			newQueue.enqueue(4)

			if newQueue.isEmpty() != false {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", false, newQueue.isEmpty())
			}
		})

		t.Run("Test Queue Length", func(t *testing.T) {
			var newQueue Queue
			if newQueue.len() != 0 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 0, newQueue.len())
			}

			newQueue.enqueue(3)
			newQueue.enqueue(4)
			newQueue.dequeue()
			newQueue.enqueue(22)
			newQueue.enqueue(99)
			newQueue.dequeue()
			newQueue.dequeue()

			if newQueue.len() != 1 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 1, newQueue.len())
			}

		})
	})
}

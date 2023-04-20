package queue

import (
	"container/list"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Test Container/List For Queue", func(t *testing.T) {
		listQueue := &LQueue{
			queue: list.New(),
		}

		t.Run("List Enqueue", func(t *testing.T) {
			listQueue.Enqueue("Snap")
			listQueue.Enqueue(123)
			listQueue.Enqueue(true)
			listQueue.Enqueue(212.545454)

			if listQueue.Len() != 4 {
				t.Errorf("List Enqueue is not correct expected %d but got %d", 4, listQueue.Len())
			}
		})

		t.Run("List Dequeue", func(t *testing.T) {

			err := listQueue.Dequeue()

			if err != nil {
				t.Error("got an unexpected error ", err)
			}
			if listQueue.Len() != 3 {
				t.Errorf("List Dequeue is not correct expected %d but got %d", 3, listQueue.Len())
			}
		})

		t.Run("List Front", func(t *testing.T) {

			err := listQueue.Dequeue()

			if err != nil {
				t.Error("got an unexpected error ", err)
			}

			result, err := listQueue.Front()

			if err != nil {
				t.Error("got an unexpected error ", err)
			}

			if result != true {
				t.Errorf("List Front is not correct expected %v but got %v", true, result)
			}
		})

		t.Run("List Back", func(t *testing.T) {

			err := listQueue.Dequeue()

			if err != nil {
				t.Error("got an unexpected error ", err)
			}

			result, err := listQueue.Back()

			if err != nil {
				t.Error("got an unexpected error ", err)
			}

			if result != 212.545454 {
				t.Errorf("List Back is not correct expected %v but got %v", 212.545454, result)
			}
		})

		t.Run("List Length", func(t *testing.T) {

			listQueue.Enqueue("Snap")

			err := listQueue.Dequeue()

			if err != nil {
				t.Error("got an unexpected error ", err)
			}

			if listQueue.Len() != 1 {
				t.Errorf("List Length is not correct expected %v but got %v", 1, listQueue.Len())
			}
		})

		t.Run("List Empty", func(t *testing.T) {

			err := listQueue.Dequeue()

			if err != nil {
				t.Error("got an unexpected error ", err)
			}

			if !listQueue.Empty() {
				t.Errorf("List Empty is not correct expected %v but got %v", true, listQueue.Empty())
			}
		})

	})

}

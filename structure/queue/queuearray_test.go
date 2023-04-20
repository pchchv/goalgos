package queue

import (
	"testing"
)

func TestQueueArray(t *testing.T) {
	// Handle Queue Array
	t.Run("Test Queue Array", func(t *testing.T) {
		t.Run("Test EnQueue", func(t *testing.T) {
			EnQueue(2)
			EnQueue(23)
			EnQueue(45)
			EnQueue(66)

			if FrontQueue() != 2 && BackQueue() != 66 {
				t.Errorf("Test EnQueue is wrong the result must be %v and %v but got %v and %v", 2, 66, FrontQueue(), BackQueue())
			}

		})

		t.Run("Test DeQueue", func(t *testing.T) {
			EnQueue(2)
			EnQueue(23)
			EnQueue(45)
			EnQueue(66)

			DeQueue()
			DeQueue()

			if DeQueue() != 45 {
				t.Errorf("Test DeQueue is wrong the result must be %v but got %v", 45, DeQueue())
			}
		})

		ListQueue = []any{}

		t.Run("Test Queue isEmpty", func(t *testing.T) {

			if IsEmptyQueue() != true {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", true, IsEmptyQueue())
			}

			EnQueue(3)
			EnQueue(4)

			if IsEmptyQueue() != false {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", false, IsEmptyQueue())
			}
		})

		ListQueue = []any{}
		t.Run("Test Queue Length", func(t *testing.T) {
			if LenQueue() != 0 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 0, LenQueue())
			}

			EnQueue(3)
			EnQueue(4)
			DeQueue()
			EnQueue(22)
			EnQueue(99)
			DeQueue()
			DeQueue()

			if LenQueue() != 1 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 1, LenQueue())
			}

		})

	})
}

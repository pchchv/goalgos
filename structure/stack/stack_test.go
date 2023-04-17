package stack

import (
	"reflect"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestStackArray(t *testing.T) {
	t.Run("Stack With Array", func(t *testing.T) {

		stackPush(2)
		stackPush(3)

		t.Run("Stack Push", func(t *testing.T) {
			if !reflect.DeepEqual([]any{3, 2}, stackArray) {
				t.Errorf("Stack Push is not work we expected %v but got %v", []any{3, 2}, stackArray)
			}
		})

		pop := stackPop()

		t.Run("Stack Pop", func(t *testing.T) {

			if stackLength() == 2 && pop != 3 {
				t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
			}
		})

		stackPush(2)
		stackPush(83)

		t.Run("Stack Peak", func(t *testing.T) {

			if stackPeak() != 83 {
				t.Errorf("Stack Peak is not work we expected %v but got %v", 83, stackPeak())
			}
		})

		t.Run("Stack Length", func(t *testing.T) {
			if stackLength() != 3 {
				t.Errorf("Stack Length is not work we expected %v but got %v", 3, stackLength())
			}
		})

		t.Run("Stack Empty", func(t *testing.T) {
			if stackEmpty() == true {
				t.Errorf("Stack Empty is not work we expected %v but got %v", false, stackEmpty())
			}

			stackPop()
			stackPop()
			stackPop()

			if stackEmpty() == false {
				t.Errorf("Stack Empty is not work we expected %v but got %v", true, stackEmpty())
			}
		})
	})
}

// TestStackLinkedList for testing Stack with LinkedList
func TestStackLinkedList(t *testing.T) {
	var newStack Stack

	newStack.push(1)
	newStack.push(2)

	t.Run("Stack Push", func(t *testing.T) {
		result := newStack.show()
		expected := []any{2, 1}
		for x := range result {
			if result[x] != expected[x] {
				t.Errorf("Stack Push is not work, got %v but expected %v", result, expected)
			}
		}
	})

	t.Run("Stack isEmpty", func(t *testing.T) {
		if newStack.isEmpty() {
			t.Error("Stack isEmpty is returned true but expected false", newStack.isEmpty())
		}

	})

	t.Run("Stack Length", func(t *testing.T) {
		if newStack.len() != 2 {
			t.Error("Stack Length should be 2 but got", newStack.len())
		}
	})

	newStack.pop()
	pop := newStack.pop()

	t.Run("Stack Pop", func(t *testing.T) {
		if pop != 1 {
			t.Error("Stack Pop should return 1 but is returned", pop)
		}

	})

	newStack.push(52)
	newStack.push(23)
	newStack.push(99)
	t.Run("Stack Peak", func(t *testing.T) {
		if newStack.peak() != 99 {
			t.Error("Stack Peak should return 99 but got ", newStack.peak())
		}
	})
}

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

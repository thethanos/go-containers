package tests

import (
	"testing"

	"github.com/thethanos/go-containers/containers"
)

func TestStackPushInt(t *testing.T) {

	stack := containers.NewStack[int]()

	for i, value := range inputData {

		stack.Push(value)
		if stack.Top() != value {
			t.Error("failed to push value, wrong top value")
		}

		if stack.Size() != i+1 {
			t.Error("failed to push value, wrong stack size")
		}
	}

}

func TestStackPushStruct(t *testing.T) {

	stack := containers.NewStack[testType]()

	for i, value := range inputData {

		stack.Push(testType{value: value})
		if stack.Top().value != value {
			t.Error("failed to push value, wrong top value")
		}

		if stack.Size() != i+1 {
			t.Error("failed to push value, wrong stack size")
		}
	}
}

func TestStackPop(t *testing.T) {

	stack := containers.NewStack[int]()

	for _, value := range inputData {
		stack.Push(value)
	}

	for !stack.Empty() {
		stack.Pop()
	}

	if stack.Size() != 0 {
		t.Error("wrong stack size")
	}
}

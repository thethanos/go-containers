package tests

import (
	"testing"

	"github.com/thethanos/go-containers/containers"
)

const (
	inputSize = 100000
)

type testType struct {
	value int
}

func TestPushFrontInt(t *testing.T) {

	lst := containers.NewList[int]()
	for i := 0; i < inputSize; i++ {
		lst.PushFront(i)

		if lst.Front() != i {
			t.Error("Front value does not match")
		}
	}

	if lst.Size() != inputSize {
		t.Error("Size does not match")
	}
}

func TestPushFrontStruct(t *testing.T) {

	lst := containers.NewList[testType]()
	for i := 0; i < inputSize; i++ {
		lst.PushFront(testType{value: i})

		if lst.Front().value != i {
			t.Error("Front value does not match")
		}
	}

	if lst.Size() != inputSize {
		t.Error("Size does not match")
	}
}

func TestPushFrontPointer(t *testing.T) {

	lst := containers.NewList[*testType]()
	for i := 0; i < inputSize; i++ {
		lst.PushFront(&testType{value: i})

		if lst.Front().value != i {
			t.Error("Front value does not match")
		}
	}

	if lst.Size() != inputSize {
		t.Error("Size does not match")
	}

}

func TestPushBackInt(t *testing.T) {

	lst := containers.NewList[int]()
	for i := 0; i < inputSize; i++ {
		lst.PushBack(i)

		if lst.Back() != i {
			t.Error("Back value does not match")
		}
	}

	if lst.Size() != inputSize {
		t.Error("Size does not match")
	}
}

func TestPushBackStruct(t *testing.T) {

	lst := containers.NewList[testType]()
	for i := 0; i < inputSize; i++ {
		lst.PushBack(testType{value: i})

		if lst.Back().value != i {
			t.Error("Front value does not match")
		}
	}

	if lst.Size() != inputSize {
		t.Error("Size does not match")
	}
}

func TestPushBackPointer(t *testing.T) {

	lst := containers.NewList[*testType]()
	for i := 0; i < inputSize; i++ {
		lst.PushBack(&testType{value: i})

		if lst.Back().value != i {
			t.Error("Front value does not match")
		}
	}

	if lst.Size() != inputSize {
		t.Error("Size does not match")
	}
}

func TestPopFront(t *testing.T) {

	lst := containers.NewList[int]()
	for i := 0; i < inputSize; i++ {
		lst.PushFront(i)
	}

	for !lst.Empty() {

		temp := lst.Front()
		lst.PopFront()

		if !lst.Empty() && lst.Front() == temp {
			t.Error("Failed to pop front value")
		}
	}

	if lst.Empty() && lst.Size() != 0 {
		t.Error("List size is not equal to 0")
	}
}

func TestPopBack(t *testing.T) {

	lst := containers.NewList[int]()
	for i := 0; i < inputSize; i++ {
		lst.PushBack(i)
	}

	for !lst.Empty() {

		temp := lst.Back()
		lst.PopBack()

		if !lst.Empty() && lst.Back() == temp {
			t.Error("Failed to pop front value")
		}
	}

	if lst.Empty() && lst.Size() != 0 {
		t.Error("List size is not equal to 0")
	}
}

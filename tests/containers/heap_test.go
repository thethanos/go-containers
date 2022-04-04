package tests

import (
	"testing"

	"github.com/thethanos/go-containers/containers"
)

var inputData = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestHeapPushInt(t *testing.T) {

	heap := containers.NewHeap(func(i, j int) bool { return i > j })

	for i, value := range inputData {

		heap.Push(value)
		if heap.Top() != value {
			t.Error("failed to push value, wrong top value")
		}

		if heap.Size() != i+1 {
			t.Error("failed to push value, wrong heap size")
		}
	}
}

func TestHeapPushStruct(t *testing.T) {

	heap := containers.NewHeap(func(i, j testType) bool { return i.value > j.value })

	for i, value := range inputData {

		heap.Push(testType{value: value})
		if heap.Top().value != value {
			t.Error("failed to push value, wrong top value")
		}

		if heap.Size() != i+1 {
			t.Error("failed to push value, wrong heap size")
		}
	}
}

func TestHeapPushSlice(t *testing.T) {

	heap := containers.NewHeap(func(i, j int) bool { return i > j })
	heap.PushSlice(inputData)

	if heap.Size() != len(inputData) {
		t.Error("failed to push slice data, wrong heap size")
	}
}

func TestHeapPop(t *testing.T) {

	heap := containers.NewHeap(func(i, j int) bool { return i > j })
	heap.PushSlice(inputData)

	for !heap.Empty() {
		heap.Pop()
	}

	if heap.Size() != 0 {
		t.Error("wrong heap size")
	}
}

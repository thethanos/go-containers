package tests

import (
	"testing"

	"github.com/thethanos/go-containers/algorithms"
)

type testDataInt struct {
	data   []int
	value  int
	result int
}

var testDataIntArray = []testDataInt{
	{data: []int{0, 1, 2, 3, 4, 5, 6, 7}, value: 5, result: 5},
	{data: []int{11, 12, 44, 77, 245, 3}, value: 22, result: -1},
	{data: []int{}, value: 1, result: -1},
}

type testDataString struct {
	data   []string
	value  string
	result int
}

var testDataStringArray = []testDataString{
	{data: []string{"red", "green", "blue"}, value: "blue", result: 2},
	{data: []string{"red", "green", "orange"}, value: "purple", result: -1},
	{data: []string{}, value: "yellow", result: -1},
}

func TestFindInt(t *testing.T) {

	for _, data := range testDataIntArray {

		res := algorithms.Find(data.data, data.value)
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

func TestFindString(t *testing.T) {

	for _, data := range testDataStringArray {

		res := algorithms.Find(data.data, data.value)
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

func TestFindIfInt(t *testing.T) {

	for _, data := range testDataIntArray {

		res := algorithms.FindIf(data.data, func(value int) bool { return value > 4 && value < 6 })
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

func TestFindIfString(t *testing.T) {

	for _, data := range testDataStringArray {

		res := algorithms.FindIf(data.data, func(value string) bool { return len(value) == 4 })
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

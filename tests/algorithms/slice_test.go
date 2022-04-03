package tests

import (
	"testing"

	"github.com/thethanos/go-containers/algorithms"
)

type testDataFindInt struct {
	data   []int
	value  int
	result bool
}

var testDataFindIntArray = []testDataFindInt{
	{data: []int{0, 1, 2, 3, 4, 5, 6, 7}, value: 5, result: true},
	{data: []int{11, 12, 44, 77, 245, 3}, value: 22, result: false},
	{data: []int{}, value: 1, result: false},
}

func TestFindInt(t *testing.T) {

	for _, data := range testDataFindIntArray {

		res := algorithms.Find(data.data, data.value)
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

type testDataFindString struct {
	data   []string
	value  string
	result bool
}

var testDataFindStringArray = []testDataFindString{
	{data: []string{"red", "green", "blue"}, value: "blue", result: true},
	{data: []string{"red", "green", "orange"}, value: "purple", result: false},
	{data: []string{}, value: "yellow", result: false},
}

func TestFindString(t *testing.T) {

	for _, data := range testDataFindStringArray {

		res := algorithms.Find(data.data, data.value)
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

type testDataFindIfInt struct {
	data   []int
	result int
}

var testDataFindIfIntArray = []testDataFindIfInt{
	{data: []int{0, 1, 2, 3, 4, 5, 6, 7}, result: 5},
	{data: []int{11, 12, 44, 77, 245, 3}, result: -1},
	{data: []int{}, result: -1},
}

func TestFindIfInt(t *testing.T) {

	for _, data := range testDataFindIfIntArray {

		res := algorithms.FindIf(data.data, -1, func(value int) bool { return value > 4 && value < 6 })
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

type testDataFindIfString struct {
	data   []string
	value  string
	result string
}

var testDataFindIfStringArray = []testDataFindIfString{
	{data: []string{"red", "green", "blue"}, result: "blue"},
	{data: []string{"red", "green", "orange"}, result: "not found"},
	{data: []string{}, result: "not found"},
}

func TestFindIfString(t *testing.T) {

	for _, data := range testDataFindIfStringArray {

		res := algorithms.FindIf(data.data, "not found", func(value string) bool { return len(value) == 4 })
		if res != data.result {
			t.Error("result value does not match")
		}
	}
}

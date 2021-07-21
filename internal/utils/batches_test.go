package utils

import (
	"errors"
	"reflect"
	"testing"
)

type testBatches struct {
	in       []int
	size     int
	expected [][]int
	err      error
}

func TestBatchesSuccess(t *testing.T) {
	data := []testBatches{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}, nil},
		{[]int{1, 2, 3, 4, 5, 6}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}}, nil},
		{[]int{1, 2, 3, 4, 5}, 5, [][]int{{1, 2, 3, 4, 5}}, nil},
		{[]int{1, 2, 3, 4, 5}, 6, [][]int{{1, 2, 3, 4, 5}}, nil},
		{make([]int, 0), 5, make([][]int, 0), nil},
	}
	for _, i := range data {
		if result, _ := Batches(i.in, i.size); !reflect.DeepEqual(result, i.expected) {
			t.Errorf("Case %v, result %v", i, result)
		}
	}
}

func TestBatchesErrors(t *testing.T) {
	data := []testBatches{
		{[]int{1, 2, 3, 4, 5}, 0, nil, errors.New("batch size must be less than zero")},
		{[]int{1, 2, 3, 4, 5}, -1, nil, errors.New("batch size must be less than zero")},
	}
	for _, i := range data {
		if result, err := Batches(i.in, i.size); err == nil && i.err != nil {
			t.Errorf("Case %v, result %v", i, result)
		}
	}
}

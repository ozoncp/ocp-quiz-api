package utils

import (
	"errors"
	"testing"
)

type testBatches struct {
	in       []int
	size     int
	expected [][]int
	err      error
}

func TestBatches(t *testing.T) {
	data := []testBatches{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}, nil},
		{[]int{1, 2, 3, 4, 5, 6}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}}, nil},
		{[]int{1, 2, 3, 4, 5}, 5, [][]int{{1, 2, 3, 4, 5}}, nil},
		{[]int{1, 2, 3, 4, 5}, 6, [][]int{{1, 2, 3, 4, 5}}, nil},
		{make([]int, 0), 5, make([][]int, 0), nil},

		{[]int{1, 2, 3, 4, 5}, 0, nil, errors.New("batch size must be less than zero")},
		{[]int{1, 2, 3, 4, 5}, -1, nil, errors.New("batch size must be less than zero")},
	}
	for _, i := range data {
		result, err := Batches(i.in, i.size)
		if err == nil && i.err != nil {
			t.Errorf("Case %v, result %v", i, result)
		}
		if !twoDSlicesEquals(result, i.expected) {
			t.Errorf("Case %v, result %v", i, result)
		}
	}
}

func twoDSlicesEquals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !slicesEquals(a[i], b[i]) {
			return false
		}
	}
	return true
}

func slicesEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type testSlices struct {
	a        []int
	b        []int
	expected bool
}

func TestSlicesEquals(t *testing.T) {
	data := []testSlices{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1}, []int{1, 2}, false},
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1, 3}, []int{1, 2}, false},
	}
	for _, i := range data {
		result := slicesEquals(i.a, i.b)
		if result != i.expected {
			t.Errorf("Expected %v, got %v", i, result)
		}
	}
}

type twoDTestSlices struct {
	a        [][]int
	b        [][]int
	expected bool
}

func TestTwoDSlicesEquals(t *testing.T) {
	data := []twoDTestSlices{
		{[][]int{{1, 2}}, [][]int{{1, 2}}, true},
		{[][]int{{1, 2}, {1, 2}}, [][]int{{1, 2}}, false},
		{[][]int{{1, 2}, {1, 2}}, [][]int{{1, 2}, {3, 4}}, false},
		{[][]int{{1, 2}}, [][]int{{1, 3}}, false},
		{[][]int{}, [][]int{}, true},
	}
	for _, i := range data {
		result := twoDSlicesEquals(i.a, i.b)
		if result != i.expected {
			t.Errorf("Expected %v, got %v", i, result)
		}
	}
}

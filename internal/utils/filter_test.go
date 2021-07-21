package utils

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"Slice not filtered", []int{1, 2, 3}, make([]int, 0)},
		{"Slice filtered", []int{1, 2, 3, 4}, []int{4}},
		{"Have only not hardcoded", []int{4}, []int{4}},
	}
	for _, tt := range tests {
		if got := Filter(tt.args); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Filter() = %v, want %v", got, tt.want)
		}

	}
}

package utils

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  map[string]string
	output map[string]string
}

func TestSwap(t *testing.T) {
	var data = []testCase{
		{
			map[string]string{"key": "value"},
			map[string]string{"value": "key"},
		},
		{
			map[string]string{"key1": "value1", "key2": "value2"},
			map[string]string{"value1": "key1", "value2": "key2"},
		},
		{
			map[string]string{},
			map[string]string{},
		},
		{
			map[string]string{"key1": "value1", "key2": "value1"},
			map[string]string{"value1": "key2"},
		},
	}

	for _, i := range data {
		result := Swap(i.input)
		if !reflect.DeepEqual(result, i.output) {
			t.Errorf("Case: %v, result: %v", i, result)
		}
	}
}

package utils

import (
	"github.com/ozoncp/ocp-quiz-api/internal/models"
	"reflect"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	input := []models.Quiz{
		{1, 1, 1, "link_11111"},
		{2, 1, 12, "link_22222"},
		{3, 1, 13, "link_33333"},
		{4, 1, 14, "link_44444"},
		{5, 1, 15, "link_55555"},
		{6, 1, 16, "link_66666"},
		{7, 1, 17, "link_77777"},
	}
	expected := map[uint64]models.Quiz{
		1: {1, 1, 1, "link_11111"},
		2: {2, 1, 12, "link_22222"},
		3: {3, 1, 13, "link_33333"},
		4: {4, 1, 14, "link_44444"},
		5: {5, 1, 15, "link_55555"},
		6: {6, 1, 16, "link_66666"},
		7: {7, 1, 17, "link_77777"},
	}
	if res := SliceToMap(input); !reflect.DeepEqual(res, expected) {
		t.Error("Incorrect result")
	}
}

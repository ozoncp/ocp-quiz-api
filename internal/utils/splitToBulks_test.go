package utils

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

func TestSplitToBulks(t *testing.T) {
	type args struct {
		input []models.Quiz
		size  int
	}
	input := []models.Quiz{
		{1, 1, 1, "link_11111"},
		{2, 1, 12, "link_22222"},
		{3, 1, 13, "link_33333"},
		{4, 1, 14, "link_44444"},
		{5, 1, 15, "link_55555"},
		{6, 1, 16, "link_66666"},
		{7, 1, 17, "link_77777"},
	}
	tests := []struct {
		name string
		args args
		want [][]models.Quiz
	}{
		{"With tail", args{input[0:5], 2}, [][]models.Quiz{{input[0], input[1]}, {input[2], input[3]}, {input[4]}}},
		{"Without tail", args{input[0:4], 2}, [][]models.Quiz{{input[0], input[1]}, {input[2], input[3]}}},
		{"One batch", args{input[0:5], 5}, [][]models.Quiz{input[0:5]}},
		{"Batch more than input", args{input[0:5], 6}, [][]models.Quiz{input[0:5]}},
		{"Empty input", args{make([]models.Quiz, 0), 5}, make([][]models.Quiz, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := SplitToBulks(tt.args.input, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitToBulks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitToBulksErrors(t *testing.T) {
	type args struct {
		input []models.Quiz
		size  int
	}
	input := []models.Quiz{
		{1, 1, 1, "link_11111"},
		{2, 1, 12, "link_22222"},
		{3, 1, 13, "link_33333"},
		{4, 1, 14, "link_44444"},
		{5, 1, 15, "link_55555"},
		{6, 1, 16, "link_66666"},
		{7, 1, 17, "link_77777"},
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"Equal 0", args{input[0:4], 0}, ErrBatchSizeZero},
		{"Less 0", args{input[0:4], -1}, ErrBatchSizeZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := SplitToBulks(tt.args.input, tt.args.size); !errors.Is(got, tt.want) {
				t.Errorf("SplitToBulks() = %v, want %v", got, tt.want)
			}
		})
	}
}

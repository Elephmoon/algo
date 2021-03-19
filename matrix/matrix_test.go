package matrix

import (
	"reflect"
	"testing"
)

func TestErase(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{matrix: [][]int{
				{1, 2},
				{3, 0},
				{5, 6},
			}},
			want: [][]int{
				{1, 0},
				{0, 0},
				{5, 0},
			},
		},
		{
			name: "",
			args: args{matrix: [][]int{
				{1, 2, 3},
				{1, 0, 3},
			}},
			want: [][]int{
				{1, 0, 3},
				{0, 0, 0},
			},
		},
		{
			name: "",
			args: args{matrix: [][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 0, 4, 5},
				{1, 0, 3, 4, 5},
				{1, 2, 1, 4, 0},
			}},
			want: [][]int{
				{1, 0, 0, 4, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Erase(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Erase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEraseColumn(t *testing.T) {
	type args struct {
		matrix [][]int
		column int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
				column: 1,
			},
			want: [][]int{
				{1, 0, 3},
				{1, 0, 3},
				{1, 0, 3},
				{1, 0, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EraseColumn(tt.args.matrix, tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EraseColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEraseRow(t *testing.T) {
	type args struct {
		matrix [][]int
		row    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
				row: 2,
			},
			want: [][]int{
				{1, 2, 3},
				{1, 2, 3},
				{0, 0, 0},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EraseRow(tt.args.matrix, tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EraseRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

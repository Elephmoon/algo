package algorithms

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		data []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "get [1, 2, 3, 4, 5]",
			args: args{data: []int32{5, 2, 1, 3, 4}},
			want: []int32{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

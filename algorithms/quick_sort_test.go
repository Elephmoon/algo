package algorithms

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "get [1, 2, 3, 4, 5]",
			args: args{data: []int{5, 2, 1, 3, 4}},
			want: []int{1, 2, 3, 4, 5},
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

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateArray(100000)
		b.StartTimer()
		QuickSort(arr)
	}
}

func generateArray(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(1e9))
	}
	return result
}

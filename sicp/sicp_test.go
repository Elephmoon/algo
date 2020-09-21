package sicp

import "testing"

func Test_fRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n < 3",
			args: args{2},
			want: 2,
		},
		{
			name: "n > 3",
			args: args{5},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fRecursive(tt.args.n); got != tt.want {
				t.Errorf("fRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fWithIter(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n < 3",
			args: args{2},
			want: 2,
		},
		{
			name: "n > 3",
			args: args{5},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fWithIter(tt.args.n); got != tt.want {
				t.Errorf("fWithIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

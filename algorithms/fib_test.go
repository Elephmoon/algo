package algorithms

import "testing"

func TestFib(t *testing.T) {
	type args struct {
		n uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Calc 5th number",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "Calc 20th number",
			args: args{n: 20},
			want: 6765,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

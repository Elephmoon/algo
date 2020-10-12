package algorithms

import "testing"

func TestFastPower(t *testing.T) {
	type args struct {
		a uint32
		n uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "10^2 return 100",
			args: args{
				a: 10,
				n: 2,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FastPower(tt.args.a, tt.args.n); got != tt.want {
				t.Errorf("FastPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

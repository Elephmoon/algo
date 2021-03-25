package strings

import (
	"testing"
)

func TestCountSymbols(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true",
			args: args{str: "abcd"},
			want: true,
		},
		{
			name: "returns false",
			args: args{str: "acbcd"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUniqSymbols(tt.args.str); got != tt.want {
				t.Errorf("IsUniqSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTransposition(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true",
			args: args{
				str1: "abc",
				str2: "bca",
			},
			want: true,
		},
		{
			name: "returns false",
			args: args{
				str1: "abc",
				str2: "bcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTransposition(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("IsTransposition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns %20abc",
			args: args{str: " abc"},
			want: "%20abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpace(tt.args.str); got != tt.want {
				t.Errorf("ReplaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "strings are palindrome",
			args: args{str: "taco cat"},
			want: true,
		},
		{
			name: "string isn't palindrome",
			args: args{str: "abcd"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.str); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns a2b3c4h1",
			args: args{str: "aabbbcccch"},
			want: "a2b3c4h1",
		},
		{
			name: "returns abc",
			args: args{str: "abc"},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressStr(tt.args.str); got != tt.want {
				t.Errorf("CompressStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

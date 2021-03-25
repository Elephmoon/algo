package algorithms

import (
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	type fields struct {
		top   int
		stack []uint32
		size  uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint32
		wantErr bool
	}{
		{
			name: "return error",
			fields: fields{
				top:   0,
				stack: []uint32{},
				size:  0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "return 50",
			fields: fields{
				top:   0,
				stack: []uint32{50},
				size:  4,
			},
			want:    50,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				top:   tt.fields.top,
				stack: tt.fields.stack,
				size:  tt.fields.size,
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_IsEmpty(t *testing.T) {
	type fields struct {
		top   int
		stack []uint32
		size  uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty stack",
			fields: fields{
				top:   0,
				stack: nil,
				size:  0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				top:   tt.fields.top,
				stack: tt.fields.stack,
				size:  tt.fields.size,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Push(t *testing.T) {
	type fields struct {
		top   int
		stack []uint32
		size  uint32
	}
	type args struct {
		element uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   *stack
		args   args
	}{
		{
			name: "push 10 in empty stack",
			want: &stack{
				top:   1,
				stack: []uint32{10},
				size:  1,
			},
			args: args{element: 10},
		},
		{
			name: "push 10 in not empty stack",
			fields: fields{
				top:   2,
				stack: []uint32{1, 2, 3},
				size:  2,
			},
			want: &stack{
				top:   3,
				stack: []uint32{1, 2, 3, 10},
				size:  3,
			},
			args: args{element: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				top:   tt.fields.top,
				stack: tt.fields.stack,
				size:  tt.fields.size,
			}
			s.Push(tt.args.element)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Push() got = %v, want %v", s, tt.want)
			}
		})
	}
}

package algorithms

import (
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	type fields struct {
		top   int
		stack []uint32
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
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "return 50",
			fields: fields{
				top:   0,
				stack: []uint32{50},
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
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				top:   tt.fields.top,
				stack: tt.fields.stack,
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
				top:   0,
				stack: []uint32{10},
			},
			args: args{element: 10},
		},
		{
			name: "push 10 in not empty stack",
			fields: fields{
				top:   2,
				stack: []uint32{1, 2, 3},
			},
			want: &stack{
				top:   3,
				stack: []uint32{1, 2, 3, 10},
			},
			args: args{element: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				top:   tt.fields.top,
				stack: tt.fields.stack,
			}
			s.Push(tt.args.element)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Push() got = %v, want %v", s, tt.want)
			}
		})
	}
}

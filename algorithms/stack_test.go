package algorithms

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type fields struct {
		top   int
		stack []uint32
		size  uint32
	}
	type args struct {
		element uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "push 50",
			fields: fields{
				top:   2,
				stack: make([]uint32, 2, 2),
				size:  2,
			},
			args:    args{element: 50},
			wantErr: false,
		},
		{
			name: "return error",
			fields: fields{
				top:   -1,
				stack: nil,
				size:  0,
			},
			args:    args{element: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				top:   tt.fields.top,
				stack: tt.fields.stack,
				size:  tt.fields.size,
			}
			if err := s.Push(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				reflect.DeepEqual(tt.args.element, s.stack[1])
			}
		})
	}
}

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

func TestNewStack(t *testing.T) {
	type args struct {
		size uint32
	}
	tests := []struct {
		name string
		args args
		want *stack
	}{
		{
			name: "return stack",
			args: args{size: 2},
			want: &stack{
				top:   2,
				stack: make([]uint32, 2, 2),
				size:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

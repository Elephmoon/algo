package algorithms

import (
	"testing"
)

func TestNode_Min(t *testing.T) {
	type fields struct {
		value int
		left  *Node
		right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "return -3",
			fields: fields{
				value: 10,
				left: &Node{
					value: 2,
					left: &Node{
						value: 1,
						left: &Node{
							value: -3,
						},
					},
				},
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.Min(); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

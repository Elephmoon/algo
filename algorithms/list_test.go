package algorithms

import (
	"reflect"
	"testing"
)

func TestList_Search(t *testing.T) {
	type fields struct {
		value int
		next  *List
	}
	type args struct {
		item int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *List
		wantErr bool
	}{
		{
			name: "find -5",
			fields: fields{
				value: 10,
				next: &List{
					value: 100,
					next: &List{
						value: -5,
						next: &List{
							value: 2,
							next:  nil,
						},
					},
				},
			},
			args: args{-5},
			want: &List{
				value: -5,
				next: &List{
					value: 2,
					next:  nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				value: tt.fields.value,
				next:  tt.fields.next,
			}
			got, err := l.Search(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_InsertFrontValue(t *testing.T) {
	type fields struct {
		value int
		next  *List
	}
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List
	}{
		{
			name: "insert 5",
			fields: fields{
				value: 1000,
				next:  nil,
			},
			args: args{5},
			want: &List{
				value: 5,
				next: &List{
					value: 1000,
					next:  nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				value: tt.fields.value,
				next:  tt.fields.next,
			}
			if got := l.InsertFrontValue(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertFrontValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

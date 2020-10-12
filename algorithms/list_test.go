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

func TestList_predItem(t *testing.T) {
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
			name: "find pred for 100",
			fields: fields{
				value: 10,
				next: &List{
					value: 20,
					next: &List{
						value: 30,
						next: &List{
							value: 100,
							next:  nil,
						},
					},
				},
			},
			args: args{100},
			want: &List{
				value: 30,
				next: &List{
					value: 100,
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
			got := l.predItem(tt.args.item)
			if tt.wantErr {
				t.Errorf("predItem() wantErr %v", tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("predItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_DeleteItem(t *testing.T) {
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
			name: "delete between two elements",
			fields: fields{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next:  nil,
					},
				},
			},
			args: args{2},
			want: &List{
				value: 1,
				next: &List{
					value: 3,
					next:  nil,
				},
			},
			wantErr: false,
		},
		{
			name: "delete first element",
			fields: fields{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next:  nil,
					},
				},
			},
			args: args{1},
			want: &List{
				value: 2,
				next: &List{
					value: 3,
					next:  nil,
				},
			},
			wantErr: false,
		},
		{
			name: "delete last item",
			fields: fields{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next:  nil,
					},
				},
			},
			args: args{3},
			want: &List{
				value: 1,
				next: &List{
					value: 2,
					next:  nil,
				},
			},
			wantErr: false,
		},
		{
			name: "delete second element",
			fields: fields{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next: &List{
							value: 4,
							next:  nil,
						},
					},
				},
			},
			args: args{2},
			want: &List{
				value: 1,
				next: &List{
					value: 3,
					next: &List{
						value: 4,
						next:  nil,
					},
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
			got, err := l.DeleteItem(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

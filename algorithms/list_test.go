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
			got := l.prevItem(tt.args.item)
			if tt.wantErr {
				t.Errorf("prevItem() wantErr %v", tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prevItem() got = %v, want %v", got, tt.want)
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
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("DeleteItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_InsertBack(t *testing.T) {
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
		want   *List
		args   args
	}{
		{
			name: "insert item back",
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
			want: &List{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next: &List{
							value: 99,
							next:  nil,
						},
					},
				},
			},
			args: args{item: 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				value: tt.fields.value,
				next:  tt.fields.next,
			}
			l.InsertBack(tt.args.item)
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("InsertBack() got = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestList_FindKLastElement1(t *testing.T) {
	type fields struct {
		value int
		next  *List
	}
	type args struct {
		k uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List
	}{
		{
			name: "returns last element",
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
			args: args{k: 1},
			want: &List{
				value: 3,
				next:  nil,
			},
		},
		{
			name: "returns 4th element",
			fields: fields{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next: &List{
							value: 4,
							next: &List{
								value: 5,
								next: &List{
									value: 6,
									next:  nil,
								},
							},
						},
					},
				},
			},
			args: args{k: 3},
			want: &List{
				value: 4,
				next: &List{
					value: 5,
					next: &List{
						value: 6,
						next:  nil,
					},
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
			if got := l.FindKLastElement(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindKLastElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_DeleteDups(t *testing.T) {
	type fields struct {
		value int
		next  *List
	}
	tests := []struct {
		name   string
		fields fields
		want   *List
	}{
		{
			name: "remove dups",
			fields: fields{
				value: 1,
				next: &List{
					value: 1,
					next: &List{
						value: 2,
						next: &List{
							value: 3,
							next: &List{
								value: 1,
								next:  nil,
							},
						},
					},
				},
			},
			want: &List{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next:  nil,
					},
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
			l.DeleteDups()
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("DeleteDups() got = %v, want %v", l, tt.want)
			}

		})
	}
}

func TestSumNumbers(t *testing.T) {
	type args struct {
		a *List
		b *List
	}
	tests := []struct {
		name string
		args args
		want *List
	}{
		{
			name: "returns sum of 617 and 295",
			args: args{
				a: &List{
					value: 7,
					next: &List{
						value: 1,
						next: &List{
							value: 6,
							next:  nil,
						},
					},
				},
				b: &List{
					value: 5,
					next: &List{
						value: 9,
						next: &List{
							value: 2,
							next:  nil,
						},
					},
				},
			},
			want: &List{
				value: 2,
				next: &List{
					value: 1,
					next: &List{
						value: 9,
						next:  nil,
					},
				},
			},
		},
		{
			name: "returns sum of 111 and 11",
			args: args{
				a: &List{
					value: 1,
					next: &List{
						value: 1,
						next: &List{
							value: 1,
							next:  nil,
						},
					},
				},
				b: &List{
					value: 1,
					next: &List{
						value: 1,
						next:  nil,
					},
				},
			},
			want: &List{
				value: 2,
				next: &List{
					value: 2,
					next: &List{
						value: 1,
						next:  nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumNumbers(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Len(t *testing.T) {
	type fields struct {
		value int
		next  *List
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns 1",
			fields: fields{
				value: 1,
				next:  nil,
			},
			want: 1,
		},
		{
			name: "returns 3",
			fields: fields{
				value: 0,
				next: &List{
					value: 1,
					next: &List{
						value: 3,
						next:  nil,
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				value: tt.fields.value,
				next:  tt.fields.next,
			}
			if got := l.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		list *List
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "list is palindrome",
			args: args{list: &List{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 1,
						next:  nil,
					},
				},
			}},
			want: true,
		},
		{
			name: "list isn't palindrome",
			args: args{list: &List{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next:  nil,
					},
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.list); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindromeViaCustomStack(t *testing.T) {
	type args struct {
		list *List
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is palindrome",
			args: args{list: &List{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 1,
						next:  nil,
					},
				},
			}},
			want: true,
		},
		{
			name: "isn't palindrome",
			args: args{list: &List{
				value: 1,
				next: &List{
					value: 2,
					next: &List{
						value: 3,
						next:  nil,
					},
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromeViaCustomStack(tt.args.list); got != tt.want {
				t.Errorf("IsPalindromeViaCustomStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

package algorithms

import "errors"

type List struct {
	value int
	next  *List
}

func (l *List) Search(item int) (*List, error) {
	if l == nil {
		return nil, errors.New("list is empty")
	}
	if l.value == item {
		return l, nil
	}
	return l.next.Search(item)
}

func (l *List) InsertFrontValue(item int) *List {
	newItem := List{
		value: item,
		next:  l,
	}
	return &newItem
}

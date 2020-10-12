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

func (l *List) predItem(item int) (*List, error) {
	if l == nil || l.next == nil {
		return nil, errors.New("predecessor sought on nil list")
	}
	if l.next.value == item {
		return l, nil
	}
	return l.next.predItem(item)
}

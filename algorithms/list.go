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

func (l *List) DeleteItem(item int) (*List, error) {
	removableItem, err := l.Search(item)
	if err != nil {
		return nil, err
	}
	pred := l.predItem(item)
	if pred == nil {
		return removableItem.next, nil
	}
	pred.next = removableItem.next
	return l, nil
}

func (l *List) predItem(item int) *List {
	if l == nil || l.next == nil {
		return nil
	}
	if l.next.value == item {
		return l
	}
	return l.next.predItem(item)
}

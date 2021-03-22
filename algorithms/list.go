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

func (l *List) InsertBack(item int) {
	newItem := List{
		value: item,
		next:  nil,
	}
	for l.next != nil {
		l = l.next
	}
	l.next = &newItem
}

func (l *List) DeleteItem(item int) (*List, error) {
	removableItem, err := l.Search(item)
	if err != nil {
		return nil, err
	}
	prev := l.prevItem(item)
	if prev == nil {
		return removableItem.next, nil
	}
	prev.next = removableItem.next
	return l, nil
}

func (l *List) FindKLastElement(k uint) *List {
	ptr1 := l
	ptr2 := l
	for i := 0; uint(i) < k; i++ {
		if ptr1 == nil {
			return nil
		}
		ptr1 = ptr1.next
	}
	for ptr1 != nil {
		ptr1 = ptr1.next
		ptr2 = ptr2.next
	}
	return ptr2
}

func (l *List) DeleteDups() {
	items := make(map[int]bool)
	prev := &List{}
	for l != nil {
		ok, _ := items[l.value]
		if ok {
			prev.next = l.next
		} else {
			items[l.value] = true
			prev = l
		}
		l = l.next
	}
}

func (l *List) prevItem(item int) *List {
	if l == nil || l.next == nil {
		return nil
	}
	if l.next.value == item {
		return l
	}
	return l.next.prevItem(item)
}

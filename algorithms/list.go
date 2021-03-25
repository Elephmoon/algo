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
		ok := items[l.value]
		if ok {
			prev.next = l.next
		} else {
			items[l.value] = true
			prev = l
		}
		l = l.next
	}
}

func (l *List) Len() int {
	result := 0
	for l != nil {
		result++
		l = l.next
	}
	return result
}

func SumNumbers(a, b *List) *List {
	overflow := 0
	var result *List
	var resultEnd *List
	if a.Len() > b.Len() {
		b = addZeroes(b, a.Len()-b.Len())
	} else {
		a = addZeroes(a, b.Len()-a.Len())
	}
	for b != nil && a != nil {
		resultVal := a.value + b.value + overflow
		if resultVal > 10 {
			resultVal = resultVal % 10
			overflow = 1
		}
		if result == nil {
			result = &List{
				value: resultVal,
				next:  nil,
			}
			resultEnd = result
		} else {
			resultEnd.next = &List{
				value: resultVal,
				next:  nil,
			}
			resultEnd = resultEnd.next
		}
		a = a.next
		b = b.next
	}

	return result
}

func IsPalindrome(list *List) bool {
	fast := list
	slow := list
	// this method needs the stack data structure
	var elements []int
	for fast != nil && fast.next != nil {
		elements = append(elements, slow.value)
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	position := 0
	for slow != nil {
		if elements == nil {
			return false
		}
		top := elements[position : position+1]
		if slow.value != top[0] {
			return false
		}
		slow = slow.next
		position++
	}
	return true
}

func IsPalindromeViaCustomStack(list *List) bool {
	fast := list
	slow := list
	stack := NewStack()
	for fast != nil && fast.next != nil {
		stack.Push(uint32(slow.value))
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	for slow != nil {
		top, err := stack.Pop()
		if err != nil {
			return false
		}
		if uint32(slow.value) != top {
			return false
		}
		slow = slow.next
	}

	return true
}

func addZeroes(list *List, count int) *List {
	for i := 0; i < count; i++ {
		list.InsertBack(0)
	}
	return list
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

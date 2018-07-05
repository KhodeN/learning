package linked_list

type linkedListItem struct {
	next     *linkedListItem
	previous *linkedListItem
	value    int
}

type LinkedList struct {
	head    *linkedListItem
	tail    *linkedListItem
	current *linkedListItem
}

func (l *LinkedList) Add(value int) {
	v := &linkedListItem{value: value, previous: l.tail}
	if l.head == nil {
		l.head = v
	}
	if l.tail != nil {
		l.tail.next = v
	}

	l.tail = v
}

func (l *LinkedList) Next() bool {
	if l.current != nil {
		l.current = l.current.next
	} else {
		l.current = l.head
	}

	return l.current != nil
}

func (l *LinkedList) Current() int {
	return l.current.value
}

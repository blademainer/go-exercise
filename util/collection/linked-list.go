package collection

func New[T Element]() *List[T] {
	header := &Entry[T]{}
	header.pre = header
	header.next = header
	list := &List[T]{header, header, 0}
	return list
}

func insertBefore[T Element](entry *Entry[T], element T) {
	cur := &Entry[T]{entry.pre, entry, element}
	entry.pre.next = cur
	entry.pre = cur
}

func appendTo[T Element](entry *Entry[T], element T) {
	cur := &Entry[T]{entry, entry.next, element}
	entry.next.pre = cur
	entry.next = cur
}

func (list *List[T]) Insert(t T) {
	appendTo(list.Header, t)
	list.Size++
}

func (list *List[T]) Add(t T) {
	insertBefore(list.Tail, t)
	list.Size++
}

func (list *List[T]) First() T {
	return list.Header.next.Value
}

func (list *List[T]) Last() T {
	return list.Tail.pre.Value
}

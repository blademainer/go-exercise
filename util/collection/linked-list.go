package collection

func Init() *List {
	header := &Entry{nil, nil, nil}
	header.pre = header
	header.next = header
	list := &List{header, header, 0}
	return list
}

func insertBefore(entry *Entry, element interface{}) {
	cur := &Entry{entry.pre, entry, element}
	entry.pre.next = cur
	entry.pre = cur
}

func appendTo(entry *Entry, element interface{}) {
	cur := &Entry{entry, entry.next, element}
	entry.next.pre = cur
	entry.next = cur
}

func (list *List) Insert(t interface{}) {
	appendTo(list.Header, t)
	list.Size++
}

func (list *List) Add(t interface{}) {
	insertBefore(list.Tail, t)
	list.Size++
}

func (list *List) First() interface{} {
	return list.Header.next.Value
}

func (list *List) Last() interface{} {
	return list.Tail.pre.Value
}

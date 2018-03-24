package collection


type List struct {
	Header *Entry
	Tail *Entry
	Size int64
}

type Entry struct {
	pre   *Entry;
	next  *Entry;
	Value interface{};
}

type ListFunc interface {
	Add(t interface{})
	First(list *List) interface{}
	Last(list *List) interface{}
}
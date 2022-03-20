package collection

import (
	"fmt"
	"strings"
)

type Element interface {
	int64 | float64 | string | int
}

type List[T Element] struct {
	Header *Entry[T]
	Tail   *Entry[T]
	Size   int
}

func (list *List[T]) String() string {
	cur := list.Header
	sb := &strings.Builder{}
	delimiter := ""
	for i := 0; i < list.Size; i++ {
		cur = cur.next
		sb.WriteString(delimiter)
		sb.WriteString(fmt.Sprint(cur.Value))
		delimiter = " -> "
	}
	return sb.String()
}

type Entry[T Element] struct {
	pre   *Entry[T]
	next  *Entry[T]
	Value T
}

// type ListFunc[T] interface {
// 	Add(t T)
// 	First(list *List) T
// 	Last(list *List) T
// }

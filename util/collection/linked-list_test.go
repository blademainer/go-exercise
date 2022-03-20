package collection

import (
	"fmt"
	"testing"
)

func TestList_Insert(t *testing.T) {
	list := New[string]()
	list.Add("a2")
	list.Add("b")
	list.Add("c")
	list.Insert("a1")
	fmt.Printf("list: %s\n", list.String())
}

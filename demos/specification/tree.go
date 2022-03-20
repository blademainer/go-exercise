package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
	// if t == nil {
	//	return
	// }
	// walkImpl(t.Left, ch)
	// ch <- t.Value
	// walkImpl(t.Right, ch)
	if t.Left != nil {
		walkImpl(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkImpl(t.Right, ch)
	}

}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, err1 := <-c1
		v2, err2 := <-c2
		if !err1 || !err2 {
			return err1 == err2
		}
		if v1 != v2 {
			return false
		}

	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(100000) != tree.New(200000): ")
	if !Same(tree.New(100000), tree.New(200000)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}

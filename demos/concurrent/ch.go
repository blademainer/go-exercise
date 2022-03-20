package main

import "fmt"

type StringCh <-chan string

func UpString() StringCh {
	ches := make(chan string, 1)
	go func() {
		for s := 0; s <= 100; s++ {
			ches <- string(fmt.Sprintf("%d", s))
		}
	}()

	return ches
}

func main() {
	upString := UpString()
	for s := 0; s <= 100; s++ {
		fmt.Println(<-upString)
	}
	// for i := range upString {
	//	fmt.Println(i)
	// }
}

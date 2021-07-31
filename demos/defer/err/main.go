package main

import "fmt"

func A() (err error) {
	defer func() {
		if err != nil {
			fmt.Println("error: ", err.Error())
			// err = fmt.Errorf("rewrite err")
		}
	}()
	err2 := fmt.Errorf("returns err2")
	return err2
}

func main() {
	err := A()
	if err != nil {
		fmt.Println(err.Error())
	}
}

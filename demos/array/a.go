package main

import "fmt"

type Dog struct {
	Name string
}


func f1() (result int) {
	defer func() {
		fmt.Println("defer result is: ", result)
		result *= 7
		fmt.Println("defer after result is: ", result)
	}()
	return 6
}

func f2() (result int) {
	tmp := 5
	defer func() {
		fmt.Println("defer result is: ", result)
		fmt.Println("defer tmp is: ", tmp)
		tmp = tmp + 5
		fmt.Println("defer after result is: ", result)
		fmt.Println("defer after tmp is: ", tmp)

	}()
	return tmp
}

func main() {
	var dogs = []Dog{}
	dogs = append(dogs, Dog{Name: "alice"}, Dog{Name: "bob"})

	var copys []*Dog
	for _, d := range dogs {
		fmt.Println(&d)
		copys = append(copys, &d)
	}

	fmt.Println(copys)

	fmt.Println(f1())
	fmt.Println(f2())
}

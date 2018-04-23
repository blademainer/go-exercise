package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	age int
}


func increaseAge(person Person){
	person.age += 1
}

func increaseAgePoint(person *Person){
	person.age += 1
}

func increaseI(i int) {
	i++
}

func increaseIByPoint(i *int) {
	*i++
}

func (p *Person) String() string{
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

func main() {
	someOne := Person{name: "zhangsan", age: 1}
	increaseAge(someOne)
	fmt.Println(someOne.age) // 1
	increaseAgePoint(&someOne)
	fmt.Println(someOne.age) // 2


	i := 5
	increaseI(i)
	fmt.Println(i) // 5
	increaseIByPoint(&i)
	fmt.Println(i) // 6
}

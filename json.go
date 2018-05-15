package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Age  uint   `json:"age"`
	Name string `json:"name"`
}

func (student *Student) String() string {
	bytes, _ := json.Marshal(student)
	return string(bytes)
}

func main() {
	student := Student{Age: 18, Name: "zhang san"}
	bytes, _ := json.Marshal(student)
	fmt.Println(string(bytes))

	target := &Student{}
	json.Unmarshal(bytes, target)
	fmt.Printf("%s \n", target)
	fmt.Println(*target == student)
}

package main

import (
	"fmt"
	"github.com/blademainer/go-exercise/pkg/field"
)

func main() {
	type Person struct {
		Age  uint8  `form:"age"`
		Name string `form:"name"`
	}

	parser := &field.Parser{Tag: "form", Escape: false, GroupDelimiter: '&', PairDelimiter: '='}
	parser.Tag = "form"
	person := &Person{}
	person.Name = "张三"
	person.Age = 18

	params := make(map[string][]string)
	params["name"] = []string{"张三"}
	params["age"] = []string{"李四"}

	parser.Bind(person, params)
	fmt.Println(person)

	if bytes, e := parser.Marshal(person); e == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("error happend: ", e)
	}
}

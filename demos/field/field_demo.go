package main

import (
	"fmt"
	"github.com/blademainer/go-exercise/pkg/field"
)

func main() {
	type Person struct {
		Name string `form:"name"`
		Age  uint8  `form:"age"`
	}

	parser := &field.Parser{}
	parser.Tag = "form"
	person := &Person{"zhangsan", 18}

	params := make(map[string][]string)
	params["name"] = []string{"张三"}
	params["age"] = []string{"李四"}

	parser.Bind(person, params)
	fmt.Println(person)

	if bytes, e := field.Marshal(person); e == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(e)
	}
}

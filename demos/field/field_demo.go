package main

import (
	"fmt"
	"github.com/blademainer/go-exercise/pkg/field"
	"encoding/json"
)

func main() {
	type Person struct {
		Name string `form:"name"`
		Age  uint8  `form:"age"`
		Parent *Person `form:"parent"`
	}

	parser := field.HTTP_ENCODED_FORM_PARSER
	parser.Sort = true
	parser.Tag = "form"

	parent := &Person{}
	parent.Name = "张二"
	parent.Age = 40

	person := &Person{}
	person.Name = "张三"
	person.Age = 18
	person.Parent = parent

	//params := make(map[string][]string)
	//params["name"] = []string{"张三"}
	//params["age"] = []string{"20"}
	//
	//parser.Unmarshal(person, params)
	//fmt.Println(person)

	bytes, _ := json.Marshal(person)
	fmt.Println("json: ", string(bytes))

	for i := 0; i < 100; i++ {
		if bytes, e := parser.Marshal(person); e == nil {
			fmt.Println(string(bytes))
		} else {
			fmt.Println("error happend: ", e)
		}
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

type Context struct {
	Data map[string]interface{}
}

func (c Context) SplitString() {

}

func main() {

	// http
	jsonStr := `{
	"names": [{"name":"1"}, {"name":"2"}],
	  "age": 10
	}`
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	// response convert
	parse, err := template.New("t").Parse(
		`first name is {{ (index .names 0).name }}
second name is {{ (index .names 1).name }}
`,
	)
	if err != nil {
		panic(err)
	}
	b := &bytes.Buffer{}
	err = parse.Execute(b, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.String())
}

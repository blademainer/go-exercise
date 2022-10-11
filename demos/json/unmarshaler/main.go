package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Foo struct {
	Bar *Bar `json:"bar"`
}

type Bar struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type fakeBar struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (b *Bar) UnmarshalJSON(bytes []byte) error {
	fmt.Println(string(bytes))
	raw := strings.ReplaceAll(strings.Trim(string(bytes), `"`), `\"`, `"`)
	fmt.Println(raw)
	fb := &fakeBar{}
	err := json.Unmarshal([]byte(raw), fb)
	if err != nil {
		return err
	}
	b.Age = fb.Age
	b.Name = fb.Name
	return nil
	// return json.Unmarshal(bytes, b)
}

func main() {
	f := &Foo{Bar: &Bar{}}
	err := json.Unmarshal(
		[]byte(`
{
"bar": "{\"name\":\"zhangsan\"}"
}
`), f,
	)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println(f.Bar.Name)
}

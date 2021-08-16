package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type a struct {
	Name string
	Age  int
}

func (i *a) SetName(name string) {
	i.Name = name
}
func (i *a) SetAge(age int) error {
	i.Age = age
	return nil
}

func (i *a) Mul(a, b int) int{
	return a*b
}

func main() {
	aa := &a{
		Name: "zhangsan",
		Age:  18,
	}

	ageTmpl, err := template.New("test").Parse("{{ .SetAge (.Mul .Age 7) }}")
	if err != nil {
		panic(err.Error())
	}

	w := &bytes.Buffer{}
	err = ageTmpl.Execute(w, aa)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(aa.Age)
}

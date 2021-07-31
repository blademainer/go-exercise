package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"
)

type Key string

type KV map[Key]interface{}

func NewKey(name string) Key {
	return Key(name)
}

func main() {
	kv := KV{}
	kv["a"] = "b"
	kv["Name"] = "b"
	receive := reflect.ValueOf(kv)
	nameVal := reflect.ValueOf("Name")
	fmt.Println(receive.Type().Key())
	fmt.Println(nameVal.Type())
	fmt.Println(nameVal.Type())
	if !nameVal.Type().ConvertibleTo(receive.Type().Key()) {
		panic("not convertible")
	}
	c := nameVal.Convert(receive.Type().Key())

	if !c.Type().AssignableTo(receive.Type().Key()) {
		panic("not assignable")
	}
	index := receive.MapIndex(c)
	fmt.Println(index)

	funcMap := template.FuncMap{
		"newKey": NewKey,
	}
	tmpl, err := template.
		New("test").
		Funcs(funcMap).
		Parse(" Hello {{ .Name }}")
	if err != nil{
		log.Printf(err.Error())
	}
	err = tmpl.Execute(os.Stderr, kv)
	if err != nil {
		log.Printf("%v\n", err.Error())
	}

	tmpl1, err := template.
		New("test").
		Funcs(funcMap).
		Parse(`{{ eq "a" "b" }}`)

	if err != nil {
		log.Printf("%v\n", err.Error())
	}
	err = tmpl1.Execute(os.Stdout, "")
}

package main

import (
	"fmt"
	"reflect"
)

type Root struct {
	Child *Child
	CC    **Child
}

type Child struct {
	Name string
	Age  int
}

func main() {
	c := &Child{Name: "zhangsan"}
	r := &Root{Child: c}
	of := reflect.ValueOf(r).Elem()
	fmt.Println(of.Field(0))

	// Settings
	// pointer to struct - addressable
	ps := reflect.ValueOf(c)
	// struct
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		// exported field
		f := s.FieldByName("Name")
		if f.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if f.CanSet() {
				// change value of N
				f.Set(reflect.ValueOf("aaa"))
				fmt.Println("setting field: name")
			}
		}
	}
	// N at end
	fmt.Println("new name: " + c.Name)

	r2 := &Root{}
	rv := reflect.ValueOf(r2)
	elem := rv.Elem()
	if elem.Kind() == reflect.Struct {
		childField := elem.FieldByName("Child")
		if !childField.IsValid() {
			panic(childField)
		}
		if !childField.CanSet() {
			panic(childField)
		}
		fmt.Println("childField.Type(): ", childField.Type())
		childValue := reflect.New(childField.Type().Elem())
		childField.Set(childValue)
		ageField := childField.Elem().FieldByName("Age")
		if !ageField.IsValid() {
			panic(ageField)
		}
		ageField.Set(reflect.ValueOf(19))

		// CC
		ccField := elem.FieldByName("CC")
		t := ccField.Type()
		for t.Kind() == reflect.Ptr {
			value := reflect.New(t)
			fmt.Printf("setting value: %v to field: %v \n", value.Kind(), ccField.Kind())
			ccField.Set(value.Elem())
			t = t.Elem()
		}
	}

	fmt.Println("age: ", r2.Child.Age)
	fmt.Println("age: ", r2.CC)
}

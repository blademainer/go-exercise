package main

import (
	"reflect"
	"fmt"
	"encoding/json"
)

func printType(t reflect.Type) {
	fmt.Printf("instance: %v valueOf: %v type: %T typeOfValue: %T, kind: %v \n", t, reflect.ValueOf(t), t, reflect.TypeOf(reflect.ValueOf(t)), reflect.TypeOf(t).Kind())
}

func jsonCopy(origin interface{}, target interface{}) {
	bytes, e := json.Marshal(origin)
	fmt.Println("json error: ", e)
	json.Unmarshal(bytes, target)
}


func structCopy(origin interface{}){

}

func main() {
	type Person struct {
		Name string
		Age  uint
	}
	s := &Person{Name: "zhangsan"}
	printType(reflect.TypeOf(s))
	typ := reflect.TypeOf(reflect.TypeOf(s))
	if typ.Kind() == reflect.Ptr {
		fmt.Println("is pointer!!")
		//fmt.Printf("ele: %v  valid: %s field: %v \n ", typ, typ.String(), typ.NumField())
		printType(typ)
		typ = typ.Elem()
		typ = reflect.TypeOf(s)
		printType(typ)
	}
	v := reflect.ValueOf(s)
	vType := reflect.TypeOf(typ)
	value := reflect.New(typ)
	printType(value.Type())
	printType(value.Type().Elem())
	elem := value.Type().Elem()
	jsonCopy(s, elem)
	fmt.Println("afet copy...")
	printType(elem)

	fmt.Printf("origin: %s ele: %s ele.value: %v \n", s, elem.String(), elem.Kind())
	pointer := value.Pointer()

	i := reflect.New(reflect.ValueOf(s).Elem().Type())
	fmt.Println("From value....")
	printType(i.Type())
	convert := i.Elem().Convert(reflect.TypeOf(s).Elem())
	fmt.Printf("convert... Type: %T \n", convert)
	jsonCopy(s, convert)
	fmt.Println("convert: ", convert)

	person := i.Elem().Interface().(Person)
	i2 := reflect.New(reflect.TypeOf(person)).Elem().Convert(reflect.TypeOf(person))
	fmt.Println("i2: ", i2.String())
	fmt.Printf("origin: %s ele: %s ele.value: %v \n", i, i.Elem().String(), i.Kind())

	fmt.Printf("instance: %v valueOf: %v type: %T typeOfValue: %T copyOfType: %T pointerOfType: %v \n", s, v, typ, vType, value, pointer)
}

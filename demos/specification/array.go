package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func a(s []string) {
	fmt.Println(s)
}

func change(s *[]string) {
	a := []string{"hello", "world!"}
	*s = a
}

func InitType(i interface{}) interface{} {
	v := reflect.ValueOf(i)
	typeOf := reflect.TypeOf(i)
	value := reflect.New(typeOf)
	// reflect.Copy(v, value)
	fmt.Printf("Value: %T Type: %T new Type: %T \n", v, typeOf, reflect.TypeOf(value))
	return value
}

// dst should be a pointer to struct, src should be a struct
func Copy(dst interface{}, src interface{}) (err error) {
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		err = errors.New("dst isn't a pointer to struct")
		return
	}
	dstElem := dstValue.Elem()
	if dstElem.Kind() != reflect.Struct {
		err = errors.New("pointer doesn't point to struct")
		return
	}

	srcValue := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		err = errors.New("src isn't struct")
		return
	}

	for i := 0; i < srcType.NumField(); i++ {
		sf := srcType.Field(i)
		sv := srcValue.FieldByName(sf.Name)
		// make sure the value which in dst is valid and can set
		if dv := dstElem.FieldByName(sf.Name); dv.IsValid() && dv.CanSet() {
			dv.Set(sv)
		}
	}
	return
}

type P struct {
	Name string
	Age  int
}

func main() {

	s := []string{"1", "2", "3"}
	a(s)

	c := make([]string, 0)
	change(&c)
	fmt.Println(c)

	p := &P{Name: "张三"}
	initType := InitType(p)
	fmt.Printf("Type: %T \n", p)
	bytes, _ := json.Marshal(p)
	fmt.Println("p: ", string(bytes))
	json.Unmarshal(bytes, initType)
	fmt.Println(initType)
	b2, _ := json.Marshal(initType)
	fmt.Println("newP: ", string(b2))
}

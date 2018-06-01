package main

import (
	"reflect"
	"fmt"
	"encoding/json"
)

func main() {
	type P struct {
		Name   string
		Parent *P
	}
	p := &P{Name: "zhangsan"}

	i := CopyStruct(p)
	j, _ := json.Marshal(p)
	json.Unmarshal(j, i)
	//name := np.FieldByName("name")
	//fmt.Println("name ==== ", name)
	fmt.Printf("np==== %v type is: %v p: %v \n", i, reflect.TypeOf(i), p)

	a := []P{{Name: "zhangsan"}, {Name: "lisi"}}
	a2 := CopyStruct(a)
	j2, _ := json.Marshal(a)
	json.Unmarshal(j2, a2)
	fmt.Printf("np==== %v type is: %v p: %v \n", a2, reflect.TypeOf(a2), a)
}

func CopyStruct(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Map:
		fmt.Println("Map...")
	case reflect.Slice:
		s := reflect.ValueOf(i)
		fmt.Println("Slice...")
		for i := 0; i < s.Len(); i++ {
			fmt.Printf("Index: %v value: %v \n", i, s.Index(i))
		}
	case reflect.Chan:
		fmt.Println("Chan...")
	case reflect.Struct:
		fmt.Println("Struct...")
	case reflect.Ptr:
		t = t.Elem()
		fmt.Println("Ptr...")
		//	fv := reflect.New(ft.Type.Elem())
		//	initializeStruct(ft.Type.Elem(), fv.Elem())
		//	f.Set(fv)
	default:
		fmt.Println("Others...")
	}
	v := reflect.New(t)
	initializeStruct(t, v.Elem())
	result := v.Interface()
	return result
}

func initializeStruct(t reflect.Type, v reflect.Value) {
	if t.Kind() == reflect.Slice {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
			//case reflect.Ptr:
			//	fv := reflect.New(ft.Type.Elem())
			//	initializeStruct(ft.Type.Elem(), fv.Elem())
			//	f.Set(fv)
		default:
		}
	}
}

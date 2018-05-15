package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name   string
	Age    int
	Parent *Person
}

func main() {
	parent := Person{}
	parent.Age = 40
	parent.Name = "张三"

	p := Person{}
	p.Age = 18
	p.Name = "张四"
	p.Parent = &parent

	value := reflect.ValueOf(p)

	fields := typeFields(value.Type())
	bytes, _ := json.Marshal(fields)
	for _, f := range fields {
		fmt.Printf("name: %s, type: %v, kind: %v \n", f.Name, f.Type, f.Type.Kind())
		if f.Type.Kind() == reflect.Ptr {
			fmt.Printf("field: %s, is point! value: %v \n", f.Name, f.Type.Elem())

		}
	}

	fmt.Printf("type: %v, kind: %v, type.kind: %v \n", value.Type(), value.Kind(), string(bytes))

}

// An encodeState encodes JSON into a bytes.Buffer.
type encodeState struct {
	bytes.Buffer // accumulated output
	scratch      [64]byte
}

type encoderFunc func(e *encodeState, v reflect.Value)

type ptrEncoder struct {
	elemEnc encoderFunc
}

func (pe *ptrEncoder) encode(e *encodeState, v reflect.Value) {
	if v.IsNil() {
		e.WriteString("null")
		return
	}
	pe.elemEnc(e, v.Elem())
}

func newPtrEncoder(t reflect.Type) encoderFunc {
	enc := &ptrEncoder{typeEncoder(t.Elem())}
	return enc.encode
}

func typeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.Bool:
		return func(e *encodeState, v reflect.Value) {
			if v.Bool() {
				e.WriteString("true")
			} else {
				e.WriteString("false")
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(e *encodeState, v reflect.Value) {
			b := strconv.AppendInt(e.scratch[:0], v.Int(), 10)
			e.Write(b)
		}
	case reflect.String:
		return func(e *encodeState, v reflect.Value) {
			e.WriteString(v.String())
		}
	case reflect.Struct:
		return func(e *encodeState, v reflect.Value) {

		}
	case reflect.Ptr:
		return newPtrEncoder(t)
	default:
		return func(e *encodeState, v reflect.Value) {
			e.WriteString("null")
		}
	}
}

// typeFields returns a list of fields that JSON should recognize for the given type.
// The algorithm is breadth-first search over the set of structs to include - the top struct
// and then any reachable anonymous structs.
func typeFields(t reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i))
	}
	return fields
}

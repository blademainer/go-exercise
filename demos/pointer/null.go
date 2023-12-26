package main

import (
	"encoding/json"
)

func TestPointer(ptr interface{}) error {
	if ptr == nil {
		panic("ptr is nil")
		return nil
	}
	s, ok := ptr.(*string)
	if ok {
		if s == nil {
			panic("s is nil")
		}
		*s = ""
		return nil
	}

	structPtr, ok := ptr.(*json.Decoder)
	if ok {
		s := &json.Decoder{}
		*structPtr = *s
	}
	return nil
}

func main() {
	var s2 *string
	if err := TestPointer(s2); err != nil {
		panic(err.Error())
	}

	var s *json.Decoder
	if err := TestPointer(s); err != nil {
		panic(err.Error())
	}

}

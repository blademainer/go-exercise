package generic

import (
	"fmt"
	"reflect"
)

type Number interface {
	~string | ~int64 | ~float64 | ~int
}

type IntType int

// SumIntOrFloat sum of map
func SumIntOrFloat[K comparable, V int | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumber sum of map
func SumNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Key interface {
	~string | ~int32 | ~int64 | ~int
}

type Server[T Number] struct {
}

// NewServer create server
func NewServer[T Number]() *Server[T] {
	return &Server[T]{}
}

func (s *Server[T]) PrintKey(key T) {
	fmt.Println(reflect.TypeOf(key))
}

func (s *Server[T]) GetKeys(key ...interface{}) []interface{} {
	return []interface{}{}
}

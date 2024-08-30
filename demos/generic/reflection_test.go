package generic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectType(t *testing.T) {
	ReflectType(&Server[int]{})
	nt := NewT[string]()
	fmt.Println(reflect.TypeOf(nt))
}

func BenchmarkNewT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewT[int]()
	}
}

func BenchmarkReflectNewT(b *testing.B) {
	t := reflect.TypeOf(0)
	for i := 0; i < b.N; i++ {
		reflect.New(t)
	}
}

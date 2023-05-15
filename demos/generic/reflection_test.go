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

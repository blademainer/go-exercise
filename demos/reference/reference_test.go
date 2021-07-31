package reference

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	type p struct {
		name string
	}

	var vp *p

	var f = func(p1 *p) {
		p1 = &p{name: "test"}
	}
	f(vp)
	fmt.Println(vp)
	var f2 = func(p1 **p) {
		*p1 = &p{name: "test"}
	}
	f2(&vp)
	fmt.Println(vp)
}

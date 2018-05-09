package field

import (
	"testing"
	"fmt"
)

type Person struct {
	Name string `form:"name"`
	Age  uint8  `form:"age"`
}

func TestBind(t *testing.T) {
	parser := &Parser{Tag: "form", Escape: false, GroupDelimiter: '&', PairDelimiter: '='}
	parser.Tag = "form"
	person := &Person{"zhangsan", 18}
	params := make(map[string][]string)
	data := parser.Bind(person, params)
	fmt.Println(data)

}

func TestEncode(t *testing.T) {
	parser := &Parser{Tag: "form", Escape: false, GroupDelimiter: '&', PairDelimiter: '='}
	parser.Tag = "form"

	type Person struct {
		Name string `form:"name"`
		Age  uint8  `form:"age"`
	}

	person := &Person{"张三", 18}

	if b, e := parser.Marshal(person); e == nil {
		fmt.Println(string(b))
	} else {
		t.Fail()
	}

	m := map[string]string{"a": "b", "你好": "呵呵"}

	if b, e := parser.Marshal(m); e == nil {
		fmt.Println(string(b))
	} else {
		t.Fail()
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

type InterfaceType struct {
	A string      `json:"a"`
	B interface{} `json:"b"`
}

type MapType struct {
	A string                 `json:"a"`
	B map[string]interface{} `json:"b"`
}

type ClearStringType struct {
	A string `json:"a"`
	B string `json:"b"`
}

type ClearType struct {
	A string `json:"a"`
	B C      `json:"b"`
}

type C struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UnmarshalToMapType(jsonString string) *MapType {
	i2 := &MapType{}
	if e := json.Unmarshal([]byte(jsonString), i2); e != nil {
		fmt.Println(e)
		return nil
	} else {
		return i2
	}
}

func (m *MapType) ConvertMapTypeTo(i interface{}) (e error) {
	if bytes, e := json.Marshal(m); e != nil {
		return e
	} else {
		return json.Unmarshal(bytes, i)
	}
}

func main() {
	// m := make(map[string]string)
	// m["cc"] = "asdf"
	// m["d"] = "a"
	// m["z"] = "d"
	m := &C{Name: "zhangsan", Age: 18}
	interfaceType := InterfaceType{"a", m}
	bytes, e := json.Marshal(interfaceType)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(string(bytes))
	}

	i := &InterfaceType{}
	e = json.Unmarshal(bytes, i)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(i)
		if clearType, ok := i.B.(ClearType); ok {
			fmt.Println("ClearType: ", clearType)
		} else {
			fmt.Println("Failed to convert interface type to clearType!")
		}
	}

	i2 := &InterfaceType{}
	e = json.Unmarshal(bytes, i2)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(i)
		if clearType, ok := i.B.(ClearStringType); ok {
			fmt.Println("ClearType: ", clearType)
		} else {
			fmt.Println("Failed to convert interface type to clearType!")
		}
	}

	mapType := UnmarshalToMapType(string(bytes))
	fmt.Println(mapType)
	clearType := &ClearType{}
	mapType.ConvertMapTypeTo(clearType)
	fmt.Println("clearType: ", clearType)
}

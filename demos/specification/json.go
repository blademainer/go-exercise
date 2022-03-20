package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	type Person struct {
		Age    uint    `json:"age"`
		Name   string  `json:"name"`
		Parent *Person `json:"parent"`
	}

	student := Person{Age: 18, Name: "zhang san"}
	student.Parent = &Person{Age: 28, Name: "zhang er"}
	bytes, _ := json.Marshal(student)
	fmt.Println(string(bytes))

	// pretty json
	fmt.Println("Pretty json: ", string(PrettyJson(bytes)))
	fmt.Println("Pretty string: ", string(PrettyJson([]byte("http://baidu.com"))))

	target := &Person{}
	json.Unmarshal(bytes, target)
	fmt.Printf("%s \n", target)
	fmt.Println(*target == student)

	j := "[{\"id\":1,\"no\":\"000001\",\"ipDescribe\":\"测试\"},{\"id\":1,\"no\":\"000002\",\"ipDescribe\":\"测试2\"}]"
	fmt.Println(string(PrettyJson([]byte(j))))
}

func PrettyJson(bytes []byte) []byte {
	jsonString := strings.TrimSpace(string(bytes))

	isObject := strings.HasPrefix(jsonString, "{")
	isArray := strings.HasPrefix(jsonString, "[")

	// is json?
	if !isObject && !isArray {
		return bytes
	}

	if isArray {
		s := make([]map[string]interface{}, 1)
		ptr := &s
		json.Unmarshal(bytes, ptr)
		fmt.Println("map: ", ptr)
		pretty, _ := json.MarshalIndent(ptr, "", "    ")
		return pretty
	} else {
		m := make(map[string]interface{})
		ptr := &m
		json.Unmarshal(bytes, ptr)
		fmt.Println("map: ", ptr)
		pretty, _ := json.MarshalIndent(ptr, "", "    ")
		return pretty
	}
}

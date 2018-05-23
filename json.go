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
}

func PrettyJson(bytes []byte) []byte {
	jsonString := strings.TrimSpace(string(bytes))
	// is json?
	if !strings.HasPrefix(jsonString, "{") && !strings.HasPrefix(jsonString, "[") {
		return bytes
	}
	m := make(map[string]interface{})
	ptr := &m
	json.Unmarshal(bytes, ptr)
	fmt.Println("map: ", ptr)
	pretty, _ := json.MarshalIndent(ptr, "", "    ")
	return pretty
}

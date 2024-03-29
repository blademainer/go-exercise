package main

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string
var tmpl *template.Template

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/testdata directory in the user's GOPATH.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}

type Person struct {
	Name  string
	Age   int
	Inner string
}

func (p Person) Say(admin string) string {
	return fmt.Sprintf("Hello %v, my name is %v", admin, p.Name)
}

func main() {
	// doc: https://golang.org/pkg/text/template/
	Init()
	execute := Execute()
	fmt.Println(execute)
}

func Init() {
	abs := Path("./test.gohtml")
	var err error
	tmpl, err = template.
		New("test.gohtml").
		Funcs(
			template.FuncMap{
				"multiAge": func(person Person) int {
					return person.Age * 2
				},
			},
		).
		Option("missingkey=default").
		ParseFiles(abs)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(tmpl)
}

func Execute() string {
	p := Person{Name: "Zhangsan", Age: 11, Inner: "inner"}
	data := map[string]interface{}{
		"person": p,
		"m": map[string]interface{}{
			"name": "zhangsan",
			"age":  11,
			"say":  p.Say,
		},
	}
	return ExecuteData(data)
}

func ExecuteData(data interface{}) string {
	bf := &bytes.Buffer{}
	err := tmpl.Execute(bf, data)
	if err != nil {
		panic(err.Error())
	}
	return bf.String()
}

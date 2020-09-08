package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string

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
	Name string
	Age  int
}

func main() {
	abs := Path("./test.gohtml")

	tpl, err := template.
		New("test.gohtml").
		Funcs(template.FuncMap{
			"hello": func() string {
				return "test func"
			},
		}).
		ParseFiles(abs)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(tpl)
	tpl.Execute(os.Stdout, Person{"Zhangsan", 11})
}

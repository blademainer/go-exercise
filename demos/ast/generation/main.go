package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
	"reflect"
	"strings"
)

type A interface {
	Hello()
}

func main() {
	p, err2 := build.Import("golang.org/x/tools/go/packages", "", build.IgnoreVendor)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	//fmt.Println(p)
	fmt.Printf("p.GoFiles: %v\n", p.GoFiles)
	fmt.Printf("p.SFiles: %v\n", p.SFiles)

	cfg := &packages.Config{
		Mode: packages.LoadSyntax,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join([]string{}, " "))},
	}
	pkgs, err := packages.Load(cfg, "github.com/blademainer/go-exercise/demos/ast/generation")
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	pkg := pkgs[0]
	fmt.Println("pkg.Name: ", pkg.Name)
	fmt.Println("pkg.ID: ", pkg.ID)
	fmt.Println("pkg.CompiledGoFiles: ", pkg.CompiledGoFiles)
	fmt.Println("pkg.Imports: ", pkg.Imports)
	fmt.Println("pkg.Types: ", pkg.Types)
	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			fmt.Printf("type: %v node: %v\n", reflect.TypeOf(node), node)
			switch node.(type) {
			case *ast.GenDecl:
				decl := node.(*ast.GenDecl)
				fmt.Printf("specs %#v \n", decl.Specs)
				fmt.Printf("tok %#v \n", decl.Tok)
				switch decl.Tok {
				case token.TYPE:
					fmt.Printf("token: %v \n", decl)
					for _, spec := range decl.Specs {
						tspec := spec.(*ast.TypeSpec) // Guaranteed to succeed as this is CONST.
						fmt.Printf("tspec.name %v\n", tspec.Name)
						fmt.Printf("tspec.Comment %v\n", tspec.Comment)
						fmt.Printf("tspec.Doc %v\n", tspec.Doc)
						fmt.Printf("tspec.Assign %v\n", tspec.Assign)
						fmt.Printf("tspec.Type %v\n", tspec.Type)
					}

				}
			case *ast.InterfaceType:
				decl := node.(*ast.InterfaceType)
				fmt.Printf("interface: %v\n", decl.Interface)
				for _, field := range decl.Methods.List {
					fmt.Printf("field: %v\n", field)
				}
			case *ast.Ident:
				decl := node.(*ast.Ident)
				fmt.Printf("name: %v\n", decl.Name)
				fmt.Printf("obj: %v\n", decl.Obj)
			}
			return true
		})
	}
	json, err2 := pkg.MarshalJSON()
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	fmt.Println("json: ", string(json))
}


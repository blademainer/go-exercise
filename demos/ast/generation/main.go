package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
	"strings"
)

type AF func()

// A demo type
type A interface {
	Hello(name string) (A, error)
	HelloF(name string) (AF, error)
	Marshal(interface{}) ([][]byte, error)
}

type B struct {
}

func (b B) Hello(name string) (A, error) {
	panic("implement me")
}

func (b B) HelloF(name string) (AF, error) {
	panic("implement me")
}

func (b B) Marshal(i interface{}) ([][]byte, error) {
	panic("implement me")
}

func main() {
	p, err2 := build.Import("golang.org/x/tools/go/packages", "", build.IgnoreVendor)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	// fmt.Println(p)
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
		ast.Inspect(
			file, func(node ast.Node) bool {
				// fmt.Printf("type: %v node: %v\n", reflect.TypeOf(node), node)
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
							if tspec.Name.Name != "A" && tspec.Name.Name != "B" {
								continue
							}
							fmt.Printf("tspec.name %#v\n", tspec.Name)
							fmt.Printf("tspec.Comment %#v\n", tspec.Comment)
							fmt.Printf("tspec.Doc %#v\n", tspec.Doc)
							fmt.Printf("tspec.Assign %v\n", tspec.Assign)
							fmt.Printf("tspec.Type %v\n", tspec.Type)
							switch tspec.Type.(type) {
							case *ast.InterfaceType:
								itype := tspec.Type.(*ast.InterfaceType)
								for _, field := range itype.Methods.List {
									fmt.Println("field.Names: ", field.Names)
									fmt.Println("field.Type: ", field.Type)
									fmt.Println("field.Tag: ", field.Tag)
									switch field.Type.(type) {
									case *ast.FuncType:
										fType := field.Type.(*ast.FuncType)
										for i, f := range fType.Params.List {
											fmt.Printf("field: %#v\n", f)
											// fmt.Printf("field name: %#v\n", f.Names[0])
											fmt.Printf("field type: %#v\n", f.Type)
											switch f.Type.(type) {
											case *ast.ArrayType:
												at := f.Type.(*ast.ArrayType)
												switch at.Elt.(type) {
												case *ast.ArrayType:
													att := f.Type.(*ast.ArrayType)
													fmt.Printf("att: %#v\n", att)
												}
												id := at.Elt.(*ast.Ident)
												fmt.Printf("field[%d] type: []%s", i, id.Name)
											}

										}
										for _, f := range fType.Results.List {
											fmt.Printf("result: %#v\n", f)
											// fmt.Printf("result name: %#v\n", f.Names[0])
											fmt.Printf("result type: %#v\n", f.Type)
										}
										// fmt.Printf("fType.Params.List: %#v\n", fType.Params.List)
										// fmt.Printf("fType.Results: %#v\n", fType.Results)
									}
								}

							}
						}
					case token.FUNC:
						ft := node.(*ast.FuncType)
						fmt.Printf("func: %v\n", ft.Func)
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
				case *ast.FuncType:
					ft := node.(*ast.FuncType)
					fmt.Printf("func: %v\n", ft.Func)
				case *ast.FuncDecl:
					fd := node.(*ast.FuncDecl)
					fmt.Printf("FuncDecl: %#v\n", fd)
				case *ast.FuncLit:
					fl := node.(*ast.FuncLit)
					fmt.Printf("FuncLit: %#v\n", fl)
				default:
					fmt.Printf("default: %#v\n", node)
				}
				return true
			},
		)
	}
	json, err2 := pkg.MarshalJSON()
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	fmt.Println("json: ", string(json))
}

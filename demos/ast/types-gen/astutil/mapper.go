package astutil

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type FuncToken struct {
	Ident         *ast.Ident
	FuncName      string
	InArgNames    []string
	InTypes       []string
	InArgAndTypes []string
	OutTypes      []string
}

// BuildSignature 唯一标识
func (t FuncToken) BuildSignature() string {
	return fmt.Sprintf("%s(%s) (%s)", t.FuncName, buildArray(t.InTypes), buildArray(t.OutTypes))
}

func buildArray(s []string) string {
	sb := strings.Builder{}
	delimiter := ""
	for _, inType := range s {
		sb.WriteString(delimiter)
		sb.WriteString(inType)
		delimiter = ","
	}
	return sb.String()
}

var NotFoundError = fmt.Errorf("not found interface")

func ParseInterfaceTypeFuncs(decl *ast.GenDecl, interfaceType string) ([]*FuncToken, error) {
	if decl.Tok != token.TYPE {
		return nil, NotFoundError
	}
	for _, spec := range decl.Specs {
		tspec := spec.(*ast.TypeSpec) // Guaranteed to succeed as this is Type.
		itype, ok := tspec.Type.(*ast.InterfaceType)
		if !ok || tspec.Name.Name != interfaceType {
			continue
		}
		return ParseInterfaceFunc(itype)
	}
	return nil, NotFoundError
}

func ParseInterfaceFunc(itype *ast.InterfaceType) ([]*FuncToken, error) {
	ft := make([]*FuncToken, 0, len(itype.Methods.List))
	for _, field := range itype.Methods.List {
		fType, ok := field.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		code, err := BuildFuncCode(fType)
		if err != nil {
			return nil, err
		}
		code.Ident = field.Names[0]
		code.FuncName = field.Names[0].Name // settings func name
		ft = append(ft, code)
	}
	return ft, nil
}

func BuildFuncCode(fType *ast.FuncType) (token *FuncToken, err error) {
	token = &FuncToken{}
	for i, f := range fType.Params.List {
		fieldType, err := FieldType(f.Type)
		if err != nil {
			return nil, err
		}
		var argName string
		if len(f.Names) == 0 {
			argName = fmt.Sprintf("a%d", i)
			for isFieldNameExists(fType, argName) { // generate the unique arg name
				argName = fmt.Sprintf("%s%d", argName, i)
			}
		} else {
			argName = f.Names[0].Name
		}
		fieldToken := argName + " " + fieldType
		token.InArgNames = append(token.InArgNames, argName)
		token.InArgAndTypes = append(token.InArgAndTypes, fieldToken)
		token.InTypes = append(token.InTypes, fieldType)
	}

	for _, f := range fType.Results.List {
		fieldToken, err := FieldType(f.Type)
		if err != nil {
			return nil, err
		}
		token.OutTypes = append(token.OutTypes, fieldToken)
	}
	return
}

func isFieldNameExists(fType *ast.FuncType, argName string) bool {
	for _, field := range fType.Params.List {
		for _, name := range field.Names {
			if argName == name.Name {
				return true
			}
		}
	}
	for _, field := range fType.Results.List {
		for _, name := range field.Names {
			if argName == name.Name {
				return true
			}
		}
	}
	return false
}

func FieldType(node ast.Node) (string, error) {
	b := strings.Builder{}
	switch node.(type) {
	case *ast.ArrayType:
		b.WriteString("[]")
		at := node.(*ast.ArrayType)
		childType, err := FieldType(at.Elt)
		if err != nil {
			return "", err
		}
		b.WriteString(childType)
	case *ast.InterfaceType:
		return "interface{}", nil
	case *ast.Ident:
		id := node.(*ast.Ident)
		return id.Name, nil
	default:
		return "", fmt.Errorf("unknown type: %t when parse field token", node)
	}
	return b.String(), nil
}

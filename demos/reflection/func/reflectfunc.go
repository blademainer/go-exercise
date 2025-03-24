package main

import (
	"context"
	"fmt"
	"reflect"
)

func init() {
	_ = (Proto)((*HelloReq)(nil))
	_ = (Proto)((*HelloResp)(nil))
}

type Proto interface {
	fmt.Stringer
}

type Action func(context.Context, Proto) (Proto, error)

type HelloReq struct {
	Name string
}

func (r *HelloReq) String() string {
	return fmt.Sprintf("HelloReq{Name: %s}", r.Name)
}

type HelloResp struct {
	Message string
}

func (r *HelloResp) String() string {
	return fmt.Sprintf("HelloResp{Message: %s}", r.Message)
}

func helloAction(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	fmt.Printf("helloAction, name: %s\n", req.Name)
	return &HelloResp{
		Message: fmt.Sprintf("hello, %s", req.Name),
	}, nil
}

// aopFunc is a function that takes a function and returns a new function
func aopFunc(f any) {
	oldDoValue := reflect.ValueOf(f).Elem()
	//v := reflect.ValueOf(f)
	// Copy is needed in order to prevent infinite recursion after function wrapping.
	oldDoValueCopy := reflect.New(oldDoValue.Type()).Elem()
	oldDoValueCopy.Set(oldDoValue)
	fv := reflect.MakeFunc(oldDoValue.Type(), func(args []reflect.Value) []reflect.Value {
		fmt.Println("before")
		fmt.Printf("args: %v\n", args)
		req := args[1].Interface().(*HelloReq)
		req.Name = "fake" + req.Name
		result := oldDoValueCopy.Call(args)
		fmt.Println("after")
		fmt.Printf("result: %v\n", result)
		return result
	})
	oldDoValue.Set(fv)
}

func main() {
	hf := helloAction
	aopFunc(&hf)

	resp, err := hf(context.Background(), &HelloReq{Name: "world"})
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("resp: %s\n", resp)
}

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
func aopFunc(f any) reflect.Value {
	v := reflect.ValueOf(f)
	fv := reflect.MakeFunc(v.Type(), func(args []reflect.Value) []reflect.Value {
		fmt.Println("before")
		fmt.Printf("args: %v\n", args)
		req := args[1].Interface().(*HelloReq)
		req.Name = "fake" + req.Name
		result := v.Call(args)
		fmt.Println("after")
		fmt.Printf("result: %v\n", result)
		return result
	})
	return fv
}

func main() {
	helloActionFunc := aopFunc(helloAction)
	args := []reflect.Value{
		reflect.ValueOf(context.Background()),
		reflect.ValueOf(&HelloReq{Name: "world"}),
	}
	call := helloActionFunc.Call(args)
	if len(call) != 2 {
		fmt.Println("call length is not 2")
		return
	}
	//helloActionFunc(context.Background(), &HelloReq{Name: "world"})
}

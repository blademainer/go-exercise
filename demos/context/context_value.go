package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value") // want `should assign the result of context.WithValue to a variable`
	// ctx = context.WithValue(ctx, "nil", nil)     // want `should assign the result of context.WithValue to a variable`
	if v := ctx.Value("key").(interface{}); v != nil {
		fmt.Println(v)
	}
	if v := ctx.Value("nil"); v != nil {
		fmt.Println(v)
	} else {
		fmt.Println("nil")
	}
}

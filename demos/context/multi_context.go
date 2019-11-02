package main

import (
	"context"
	"fmt"
	"time"
)

const contextName = "name"

func Child(ctx context.Context) {
	tick := time.Tick(1 * time.Second)
	Grandson(ctx)
	for {
		select {
		case t := <-tick:
			fmt.Println("tick...", t.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Println("child done!!! but continue", ctx.Value(contextName), ctx.Err().Error())
			return
		}

	}
}

func Grandson(ctx context.Context) {
	tick := time.Tick(1 * time.Second)
	go func() {
		for {
			select {
			case t := <-tick:
				fmt.Println("tick...", t.Format(time.RFC3339))
			case <-ctx.Done():
				fmt.Println("Grandson done!!!", ctx.Value(contextName), ctx.Err().Error())
				return
			}

		}
	}()
}

func main() {
	kv := context.WithValue(context.TODO(), contextName, "todo")
	c, cancelFunc := context.WithTimeout(kv, 3*time.Second)
	defer cancelFunc()
	Child(c)

	kv2 := context.WithValue(context.Background(), contextName, "backgroud")
	c2, cancelFunc2 := context.WithTimeout(kv2, 3*time.Second)
	defer cancelFunc2()
	Child(c2)
}

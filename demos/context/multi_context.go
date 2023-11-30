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
			fmt.Println("child tick...", t.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Println("child done!!!", ctx.Value(contextName), ctx.Err().Error())
			return
		}

	}
}

func Grandson(ctx context.Context) {
	nc, cf := context.WithTimeout(ctx, 2*time.Second)
	tick := time.Tick(1 * time.Second)
	go func() {
		defer cf()
		for {
			select {
			case t := <-tick:
				fmt.Println("grandson tick...", t.Format(time.RFC3339))
			case <-nc.Done():
				fmt.Println("Grandson done!!!", nc.Value(contextName), nc.Err().Error(), time.Now().Format(time.RFC3339))
				return
			}

		}
	}()
}

func main() {
	fmt.Println("start...", time.Now().Format(time.RFC3339))
	kv := context.WithValue(context.TODO(), contextName, "todo")
	c, cancelFunc := context.WithTimeout(kv, 5*time.Second)
	defer cancelFunc()
	Child(c)

	//kv2 := context.WithValue(context.Background(), contextName, "backgroud")
	//c2, cancelFunc2 := context.WithTimeout(kv2, 3*time.Second)
	//defer cancelFunc2()
	//Child(c2)
}

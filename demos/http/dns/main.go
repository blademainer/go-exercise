package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	lookup("www.google.com")
}

func lookup(domain string) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	resolver := &net.Resolver{}
	addr, err := resolver.LookupIPAddr(ctx, domain)
	if err != nil {
		panic(err.Error())
	}
	for _, a := range addr {
		fmt.Printf("%s -> %s\n", domain, a.String())
	}
}

package main

import "net"

type CacheResolver struct {
	net.Resolver
}



func main() {
	net.DefaultResolver = &CacheResolver{}
}

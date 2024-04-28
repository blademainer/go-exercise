package main

import (
	"fmt"
	"strings"
)

func main() {
	m := statis()
	for s, i := range m {
		fmt.Printf("%s: %v\n", s, i)
	}
}

var data = []string{
	"a=b", "a=d", "b=c", "b=d", "c=d", "c=e", "d=e", "d=f", "e=f", "e=g", "f=g", "f=h", "g=h", "g=i", "h=i", "h=j",
	"i=j", "i=k", "j=k", "j=l", "k=l", "k=m", "l=m", "l=n", "m=n", "m=o", "n=o", "n=p", "o=p", "o=q", "p=q", "p=r",
	"q=r", "q=s", "r=s", "r=t", "s=t", "s=u", "t=u", "t=v", "u=v", "u=w", "v=w", "v=x", "w=x", "w=y", "x=y", "x=z",
	"y=z", "y=a", "z=a", "z=b",
}

func statis() map[string][]string {
	m := make(map[string][]string)
	for _, datum := range data {
		split := strings.Split(datum, "=")
		m[split[0]] = append(m[split[0]], split[1])
	}
	return m
}

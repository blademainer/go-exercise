package main

import "fmt"

type Interface interface {
	FuncA(params string)
}

type Channel struct {
	channel string
}


type AChannel Channel

func (aChannel AChannel) FuncA(params string){
	fmt.Printf("channe: %s params: %s \n", aChannel, params)
}

func (c Channel) A(string string){

}

func main() {
	channel := AChannel{channel: "asd"}
	channel.FuncA("pppp")
}


package main

import "fmt"

type Handler interface {
	FuncA(params string)
}

type Channel struct {
	channel string
}

type AChannel struct {
}

func (aChannel AChannel) FuncA(params string) {
	fmt.Printf("channe: %s params: %s \n", aChannel, params)
}

func Start(handler Handler, channel *Channel) {
	fmt.Printf("handler: %T channel: %T \n", handler, channel)
	s := channel.channel
	handler.FuncA(s)

}

func main() {
	channel := &AChannel{}
	fmt.Printf("AChannel type: %T, Channel type: %T \n", channel, &Channel{""})
	//channel.FuncA("pppp")
	Start(channel, &Channel{""})
}

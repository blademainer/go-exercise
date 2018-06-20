package main

import "fmt"

type AConfig struct {
	Name    string
	Age     uint8
	Address string
}

func (config *AConfig) SetName(name string) (c *AConfig) {
	config.Name = name
	c = config
	return
}

func (config *AConfig) SetAge(age uint8) (c *AConfig) {
	c = config
	config.Age = age
	return
}

func (config *AConfig) SetAddress(address string) (c *AConfig) {
	c = config
	config.Address = address
	return
}

func (*AConfig) Done(){
	fmt.Println("done config.")
}

func main() {
	c := &AConfig{}
	c.SetAddress("广东省").SetAge(18).SetName("张三").Done()
	fmt.Println("Config: ", c)
}

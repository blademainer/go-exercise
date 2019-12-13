package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/inconshreveable/go-update"
)

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		panic(err)
		// error handling
	}
	return err
}

func PrintVersion() {
	fmt.Println("v0.0.3")
}

func main() {
	args := os.Args
	fmt.Println("args:", os.Args)
	if len(args) == 1 {
		PrintVersion()
		return
	}
	PrintVersion()

	if args[1] == "update" {
		fmt.Println("ready to update!!")
		e := doUpdate("http://127.0.0.1:8080")
		if e != nil {
			fmt.Println(e.Error())
		}
		fmt.Println("updated!!")
	}
	PrintVersion()

}

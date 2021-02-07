package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r, w := io.Pipe()
	go func() {
		defer func() {
			fmt.Println("closed")
			err := w.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
		}()
		reader := bufio.NewReader(os.Stdin)

		for {
			text, _ := reader.ReadString('\n')
			if strings.TrimSpace(text) == "" {
				return
			}
			_, err := w.Write([]byte(text))
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

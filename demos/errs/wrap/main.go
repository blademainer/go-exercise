package main

import (
	"fmt"
)

type notifyError struct {
	code    string
	url     string
	message string
}

func NewNotifyError(code string, url string, message string) error {
	return &notifyError{url: url, message: message}
}

func (n *notifyError) Error() string {
	return fmt.Sprintf("code: %v url: %v error: %v", n.code, n.url, n.message)
}

func IsNotifyError(err error) bool {
	_, ok := err.(*notifyError)
	return ok
}

func main() {

}

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc(
		"/redirect", func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "https://google.com", http.StatusFound)
		},
	)
	panic(http.ListenAndServe(":8080", nil))
}

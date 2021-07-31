package tracing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	ts := &http.Transport{}
	client := NewClient(ts)
	r, err := client.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(all))

}

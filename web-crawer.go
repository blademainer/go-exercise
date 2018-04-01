package main

import (
	"fmt"
	"net/http"
	"sync"
	"regexp"
	"io/ioutil"
	url2 "net/url"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	fmt.Printf("Crawling url: [%s] \n", url)
	_, err := url2.Parse(url)
	if err != nil {
		fmt.Println("err ============ ", err.Error())
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("err2 ======= ", err)
		return
	}
	//fmt.Printf("found: %s %q\n", url, body)
	fmt.Printf("Found urls: %s by url: %s \n", urls, url)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	fetcher := OurFetcher{result: make(map[string]*Result)}
	Crawl("http://www.qq.com/", 4, fetcher)
	//time.Sleep(time.Minute)
}

// fetcher is Fetcher that returns canned results.
type OurFetcher struct {
	result map[string]*Result
	mutex  sync.Mutex
}

type Result struct {
	body string
	urls []string
}

func (fetcher OurFetcher) put(url string, result *Result) {
	fetcher.mutex.Lock()
	fetcher.result[url] = result
	fetcher.mutex.Unlock()
}

func (f OurFetcher) Fetch(url string) (string, []string, error) {
	fmt.Printf("fetch: %s \n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %T. message: %s \n",err.Error(), err.Error())
		return "", nil, fmt.Errorf("err: %s. not found url: %s", err.Error(), url)
	}

	httpBody, _ := ioutil.ReadAll(resp.Body)
	body := string(httpBody)

	//fmt.Printf("Read %s bytes to string: %s \n", n, s)
	urls := FindUrls(body)
	fmt.Printf("Found urls: %s \n", urls)
	result := &Result{body: body, urls: urls}

	f.put(url, result)

	return result.body, result.urls, nil
//	if res, ok := f.result[url]; ok {
//	}
//	return "", nil, fmt.Errorf("not found: %s", url)
}


func FindUrls(string string) []string {
	compile, e := regexp.Compile("\"http[s]://[\\w\\W]+?\"")
	if e != nil{
		panic(e)
	}

	allString := compile.FindAllString(string, -1)
	return allString
}



//func (f fetcher) Fetch(url string) (string, []string, error) {
//	if res, ok := f[url]; ok {
//		return res.body, res.urls, nil
//	}
//	return "", nil, fmt.Errorf("not found: %s", url)
//}

// fetcher is a populated fetcher.
//var fetcher = fetcher{
//	"http://golang.org/": &Result{
//		"The Go Programming Language",
//		[]string{
//			"http://golang.org/pkg/",
//			"http://golang.org/cmd/",
//		},
//	},
//	"http://golang.org/pkg/": &Result{
//		"Packages",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/cmd/",
//			"http://golang.org/pkg/fmt/",
//			"http://golang.org/pkg/os/",
//		},
//	},
//	"http://golang.org/pkg/fmt/": &Result{
//		"Package fmt",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//	"http://golang.org/pkg/os/": &Result{
//		"Package os",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//}

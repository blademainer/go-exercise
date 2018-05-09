package main

import (
	"fmt"
	"net/http"
	"sync"
	"regexp"
	"io/ioutil"
	url2 "net/url"
	"strings"
	"os"
	"hash/fnv"
	"time"
	"log"
	"gopkg.in/yaml.v2"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher OurFetcher) {
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
	done := make(chan bool)
	fmt.Printf("Found urls: %s by url: %s \n", urls, url)
	for _, u := range urls {
		go func(url string) {
			if fetcher.result[url] != nil {
				fmt.Printf("Found circulation url: %v \n", url)
				return
			}
			fmt.Printf("depth: %v \n", depth)
			Crawl(u, depth-1, fetcher)
			done <- true
		}(u)
	}

	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done
		fmt.Println("done.")
	}

	return
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

var client = &http.Client{Timeout: 10 * time.Second}

func (f OurFetcher) Fetch(url string) (string, []string, error) {
	fmt.Printf("fetch: %s \n", url)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("error: %T. message: %s \n", err.Error(), err.Error())
		return "", nil, fmt.Errorf("err: %s. not found url: %s", err.Error(), url)
	}

	httpBody, _ := ioutil.ReadAll(resp.Body)
	body := string(httpBody)

	//fmt.Printf("Read %s bytes to string: %s \n", n, s)
	urls := FindUrls(body)
	fmt.Printf("Found urls: %s \n", urls)
	result := &Result{body: body, urls: urls}

	f.put(url, result)
	write(url, "")

	return result.body, result.urls, nil
	//	if res, ok := f.result[url]; ok {
	//	}
	//	return "", nil, fmt.Errorf("not found: %s", url)
}

var dir = "html-out"

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func write(name string, body string) {
	_ = os.Mkdir(dir, os.ModePerm)
	var i uint = 1
	fmt.Println(string(i))

	path := dir + "/" + strings.Replace(name, "/", "", -1)
	fmt.Printf("Write %s as path: %s \n", name, path)
	file, e := os.Create(path)
	if e != nil {
		fmt.Println(e)
	}
	ioutil.WriteFile(path, []byte(body), os.ModePerm)
	file.Close()
}

func FindUrls(string string) []string {
	compile, e := regexp.Compile("\"http[s]://[\\w\\W]+?\"")
	if e != nil {
		panic(e)
	}

	allString := compile.FindAllString(string, -1)
	//length := len(allString)
	//var newUrls [10]string
	for i, url := range allString {
		newString := strings.Replace(url, "\"", "", -1)
		allString[i] = newString
	}
	return allString
}

type conf struct {
	StorageDir string `yaml:"storage_dir"`
	Url string `yaml:"url"`
	Depth int `yaml:"depth"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("web-crawler.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	c := conf{}
	conf := c.getConf()
	fmt.Printf("conf: %v \n", *conf)
	var url = conf.Url
	var depth = conf.Depth
	dir = conf.StorageDir
	//args := os.Args
	//if len(args) >= 3 {
	//	url = args[1]
	//	i, _ := strconv.Atoi(args[2])
	//	depth = i
	//}

	for {
		fetcher := OurFetcher{result: make(map[string]*Result)}
		Crawl(url, depth, fetcher)
	}
	//time.Sleep(time.Minute)
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

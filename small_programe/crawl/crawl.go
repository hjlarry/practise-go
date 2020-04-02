package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// 一、顺序抓取
func Serial(url string, fetcher Fetcher, fetched map[string]bool) {
	if fetched[url] {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	fetched[url] = true
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Serial(u, fetcher, fetched)
	}
	return
}

type FetchState struct {
	sync.Mutex
	fetched map[string]bool
}

// 二、用锁并发抓取
func ConcurrentMutex(url string, fetcher Fetcher, f *FetchState) {

	f.Lock()
	already := f.fetched[url]
	f.fetched[url] = true
	f.Unlock()

	if already {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			ConcurrentMutex(u, fetcher, f)
		}(u)
	}
	wg.Wait()
	return
}

func worker(url string, ch chan []string, fetcher Fetcher) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- []string{}
	} else {
		fmt.Printf("found: %s %q\n", url, body)
		ch <- urls
	}

}

func master(ch chan []string, fetcher Fetcher) {
	// n用来计数，创建一个worker+1，每个worker只会向ch里加一个item，那么当n为0的时候就是所有worker工作完的时候
	n := 1
	fetched := make(map[string]bool)
	for urls := range ch {
		for _, u := range urls {
			if fetched[u] {
				continue
			}
			fetched[u] = true
			n += 1
			go worker(u, ch, fetcher)
		}
		n -= 1
		if n == 0 {
			break
		}
	}
}

// 三、用channel并发抓取
func ConcurrentChan(url string, fetcher Fetcher) {
	ch := make(chan []string)
	go func() {
		ch <- []string{url}
	}()
	master(ch, fetcher)
}

func main() {
	Serial("https://golang.org/", fetcher, make(map[string]bool))
	fmt.Println("-------------------")
	ConcurrentMutex("https://golang.org/", fetcher, &FetchState{
		fetched: make(map[string]bool),
	})
	fmt.Println("-------------------")
	ConcurrentChan("https://golang.org/", fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

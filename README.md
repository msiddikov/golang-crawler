# Very simple and lightweight URL golang crawler 

## Installation

```go
  go get github.com/msiddikov/crawler
```

## crawler.Start
Crawls the given URL recursive through internal links and returns []crawler.UrlInfo <br>
Usage:
```go
  var res []crawler.UrlInfo
	res := crawler.Start(url, 5, 100)
```
**crawler.Start**(url string, depth int, threadsNumber int) []UrlInfo <br>
  **url:**  entry point <br>
  **depth:** - depth of recursion, set -1 to crawl through all links <br>
  **threadsNumber:** - number of concurrent treads, set -1 to run without limitation <br>
### Example:
```go
package main

import (
	"fmt"
	"time"

	"github.com/msiddikov/crawler"
)

func main() {
	url := "https://go.dev"
	started := time.Now()
	fmt.Printf("Started crawling %s \n", url)
	res := crawler.Start(url, 5, 100)
	var valid, invalid uint
	for _, v := range res {
		if v.IsValid {
			fmt.Printf("[VALID] url: %s \n", v.Url)
			valid += 1
		} else {
			fmt.Printf("[INVALID] url: %s error: %s\n", v.Url, v.Err)
			invalid += 1
		}

	}
	d := time.Now().Sub(started)
	fmt.Printf("Found %v links in : %v seconds \nValid links: %v \nBroken links: %v", len(res), d.Seconds(), valid, invalid)
}
```

### type crawler.UrlInfo


```go
	UrlInfo struct {
		Url        string
		Origin     string
		IsValid    bool
		Err        error
		Title      string
		IsExternal bool
	}
```



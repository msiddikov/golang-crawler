package crawler

import (
	"net/http"
	"sync"
	"time"
)

const timeOut = 5 * time.Second

// Crawls the given URL recursive through internal links and returns []crawler.UrlInfo
//  url:  entry point
//  depth: - depth of recursion, set -1 to crawl through all links
//  threadsNumber: - number of concurrent treads, set -1 to run without limitation
func Start(url string, depth int, threadsNumber int) []UrlInfo {
	domain, err := getDomain(url)
	if err != nil {
		info := UrlInfo{
			Url:     url,
			Err:     err,
			IsValid: false,
		}
		return []UrlInfo{info}
	}

	controlCh := make(chan struct{}, 10000)
	for i := 0; i < threadsNumber; i++ {
		controlCh <- struct{}{}
	}

	doneCh := make(chan bool)

	resCh := make(chan UrlInfo)

	go func() {
		for i := range doneCh {
			_ = i
			controlCh <- struct{}{}
		}
	}()

	var wg sync.WaitGroup

	context := ctx{
		resCh:       resCh,
		doneCh:      doneCh,
		controlCh:   controlCh,
		wg:          wg,
		noConcurLim: threadsNumber < 0,
	}
	var res []UrlInfo

	context.wg.Add(1)
	go validate(url, depth, domain, "entry point", &context)

	go func() {
		context.wg.Wait()
		close(context.resCh)
		close(context.doneCh)
	}()
	for r := range context.resCh {
		res = append(res, r)
	}

	return res
}

func validate(url string, depth int, domain string, origin string, ctx *ctx) {
	info := UrlInfo{
		Url:        url,
		IsValid:    true,
		Origin:     origin,
		IsExternal: isExternal(url, domain),
	}
	fetchChildren := !info.IsExternal

	defer func() {

		ctx.resCh <- info
		ctx.doneCh <- true
		ctx.wg.Done()
	}()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		info.IsValid = false
		info.Err = err
		fetchChildren = false

	} else {
		info.Title = getTitle(resp.Body)
	}

	if depth != 0 && fetchChildren {
		links := getLinks(resp.Body)

		ctx.wg.Add(1)
		go func() {
			for _, v := range links {
				if !ctx.noConcurLim {
					<-ctx.controlCh
				}

				ctx.mu.Lock()
				if containsUrl(ctx.fetched, v) {
					ctx.doneCh <- true
					ctx.mu.Unlock()
					continue
				} else {
					ctx.fetched = append(ctx.fetched, v)
					ctx.mu.Unlock()
				}
				ctx.wg.Add(1)
				go validate(v, depth-1, domain, url, ctx)
			}
			ctx.wg.Done()
		}()

	}
}

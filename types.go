package crawler

import (
	"sync"
)

type (
	UrlInfo struct {
		Url        string
		Origin     string
		IsValid    bool
		Err        error
		Title      string
		IsExternal bool
	}

	ctx struct {
		mu          sync.Mutex
		fetched     []string
		resCh       chan UrlInfo
		controlCh   chan struct{}
		doneCh      chan bool
		logCh       chan string
		wg          sync.WaitGroup
		noConcurLim bool
	}
)

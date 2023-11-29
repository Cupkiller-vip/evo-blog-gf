package log

import (
	"sync"
)

var (
	mu  sync.Mutex
	std = getLogger(getOptions())
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = getLogger(opts)
}

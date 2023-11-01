package cleaner

import (
	"sync"
	"time"
)

var (
	videos map[string]time.Time = make(map[string]time.Time)
	locker *sync.Mutex          = &sync.Mutex{}
)

// Add or increase the video lifetime
func Push(video string) {
	locker.Lock()
	defer locker.Unlock()

	videos[video] = time.Now()
}

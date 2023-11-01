package cleaner

import (
	"context"
	"time"

	"github.com/ppcamp/movies-to-watch/streamer/config"
	log "github.com/sirupsen/logrus"
)

func Clean(force bool) {
	for v, t := range videos {
		if time.Since(t) > config.PlaylistExpire() || force {
			locker.Lock()
			delete(videos, v)
			go DeleteInDisk(v)
			locker.Unlock()
		}
	}
}

// TODO: delete playlist and chunk for v
func DeleteInDisk(v string) {
	log.WithField("video", v).Info("video deleted")
}

func Configure(ctx context.Context, d time.Duration) {
	ticker := time.NewTicker(d)

	go func(t *time.Ticker) {
		for {
			select {
			case <-t.C:
				t.Stop()
				Clean(false)

			case <-ctx.Done():
				t.Stop()
				Clean(true)
			}
		}
	}(ticker)
}

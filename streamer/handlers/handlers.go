package handlers

import (
	"strings"

	"github.com/ppcamp/movies-to-watch/common/files"
	log "github.com/sirupsen/logrus"
)

func fileExists(p, f, ext string) bool {
	playlists, err := files.GlobSuffix(p, ext)
	if err != nil {
		log.WithError(err).Warn("failed to get playlists")
		return false
	}

	for _, v := range playlists {
		if strings.Contains(v, f) {
			return true
		}
	}
	return false
}

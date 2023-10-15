package files

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// glob returns all the templates files in the templates folder.
// We need this because the filepath subdirectory glob (**/*.tmpl) only works if we don't pass the
// prefix folder (templates/**/*.tmpl), otherwise, it returns only the files under subdirectories.
//
// This function is a workaround to match all files.
func GlobSuffix(basepath string, ext ...string) ([]string, error) {
	tmpl := make([]string, 0, 10)
	err := filepath.Walk(basepath, func(path string, info os.FileInfo, err error) error {
		for _, ext := range ext {
			if !info.IsDir() && strings.HasSuffix(path, ext) {
				log.Debug("Found: ", path)
				tmpl = append(tmpl, path)
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

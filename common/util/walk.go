package util

import (
	"github.com/laamho/turbo/common"
	"io/fs"
	"path/filepath"
)

func RelativeFilePath(dir string) (files []string) {
	err := filepath.WalkDir("template", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	common.Must(err)
	return
}

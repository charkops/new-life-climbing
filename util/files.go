package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Copies a folder from src to dst
func CopyFolder(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, src, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(dst, relPath), 0755)
		} else {
			var data, err1 = ioutil.ReadFile(filepath.Join(src, relPath))
			if err1 != nil {
				return err1
			}
			return ioutil.WriteFile(filepath.Join(dst, relPath), data, 0777)
		}
	})
}

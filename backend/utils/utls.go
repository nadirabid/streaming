package utils

import (
	"path"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetAbsolutePath(relativePath string) string {
	// basePath == the directory of this file - so we gotta go up one

	if relativePath[0] == '/' {
		return relativePath
	}

	return path.Join(basePath, "..", relativePath)
}

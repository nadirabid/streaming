package utils

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	prettyjson "github.com/hokaccha/go-prettyjson"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetBasePath() string {
	return path.Join(basePath, "..")
}

func GetAbsolutePath(relativePath string) string {
	// basePath == the directory of this file - so we gotta go up one

	if relativePath[0] == '/' {
		return relativePath
	}

	return path.Join(basePath, "..", relativePath)
}

// PrettyPrint to output as json
func PrettyPrint(v interface{}) (err error) {
	b, err := prettyjson.Marshal(v)
	if err != nil {
		fmt.Printf("prettyPrint: %v\n", err)
	}

	fmt.Println(string(b))

	return
}

func PrettyStringify(v interface{}) string {
	b, err := prettyjson.Marshal(v)
	if err != nil {
		return "<COULD_NOT_STRINGIFY>"
	}

	return string(b)
}

func GetStringOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

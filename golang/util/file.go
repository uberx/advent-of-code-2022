package util

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadFile(fileLocalPath string) string {
	_, fileCallerPath, _, ok := runtime.Caller(1)
	if !ok {
		panic("Unable to find caller of util.ReadFile")
	}

	fileAbsolutePath := path.Join(path.Dir(fileCallerPath), fileLocalPath)

	content, err := os.ReadFile(fileAbsolutePath)
	if err != nil {
		panic(err)
	}

	strContent := string(content)
	return strings.TrimRight(strContent, "\n")
}

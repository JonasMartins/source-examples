package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var PATH_SEPARATOR string = fmt.Sprintf("%c", os.PathSeparator)

func GetRootDir() (*string, error) {
	var result string = ""
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			result = dir + PATH_SEPARATOR
			return &result, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return &result, nil
}

func GetFilePath(pathFromRoot *[]string) (*string, error) {
	root, err := GetRootDir()
	if err != nil {
		return nil, err
	}

	for i, x := range *pathFromRoot {
		if i < len(*pathFromRoot)-1 {
			*root += x + PATH_SEPARATOR
		} else {
			*root += x
		}
	}
	return root, nil
}

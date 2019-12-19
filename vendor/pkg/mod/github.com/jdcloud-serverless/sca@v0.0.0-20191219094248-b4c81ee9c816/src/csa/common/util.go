package common

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(file string) ([]byte, error) {
	if !FileIsExist(file) {
		return nil, fmt.Errorf("file not exist")
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

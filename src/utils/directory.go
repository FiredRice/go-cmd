package utils

import (
	"errors"
	"io/fs"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFile(filename string) (*os.File, error) {
	exist, err := PathExists(filename)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("file already exist")
	} else {
		return os.Create(filename)
	}
}

func MkdirAll(path string, perm fs.FileMode) error {
	if ok, _ := PathExists(path); !ok {
		return os.MkdirAll(path, perm)
	}
	return nil
}

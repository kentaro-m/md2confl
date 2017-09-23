package file

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

type File struct {
	Data []byte
}

func (f *File) Open(path string) error {
	ext := filepath.Ext(path)

	if ext != ".md" && ext != ".markdown" {
		return errors.New("This file is not in markdown format.")
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return errors.New("Can't open this file.")
	}

	f.Data = data

	return nil
}

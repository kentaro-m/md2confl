package file

import (
	"errors"
	"io/ioutil"
)

// File is a type that implements the interface for a file data.
type File struct {
	Data []byte
}

// Open reads a file
func (f *File) Open(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return errors.New("can't open a file")
	}

	f.Data = data

	return nil
}

package util

import (
  "path/filepath"
  "io/ioutil"
  "log"
)

func ReadFile(path string) []byte {
  ext := filepath.Ext(path)
  // TODO: Add error handling
  if ext != ".md" && ext != ".markdown" {
    err := "This file is not in markdown format."
    return []byte(err)
  }

  data, err := ioutil.ReadFile(path)

  if err != nil {
    log.Fatal(err)
  }

  return data
}

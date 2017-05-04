package util

import (
  "fmt"
  "log"

  "github.com/atotto/clipboard"
)

func Copy(text string) error {
  err := clipboard.WriteAll(text)

  if err != nil {
    log.Fatal(err)
    return err
  }

  fmt.Println("Copied to clipboard!")
  return nil
}

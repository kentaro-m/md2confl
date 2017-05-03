package confluence

import (
  "github.com/kentaro-m/blackfriday"
)

func Convert(input []byte) string {
  renderer := blackfriday.ConfluenceRenderer(0)
	extensions := 0
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

  output := blackfriday.Markdown(input, renderer, extensions)

  return string(output)
}

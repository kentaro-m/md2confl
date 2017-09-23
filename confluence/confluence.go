package confluence

import (
	"fmt"

	"github.com/kentaro-m/blackfriday"
)

type Confluence struct {
	Contents string
}

func (c *Confluence) Convert(input []byte) {
	renderer := blackfriday.ConfluenceRenderer(0)
	extensions := 0
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

	output := blackfriday.Markdown(input, renderer, extensions)
	c.Contents = string(output)
}

func (c *Confluence) Preview() {
	fmt.Println(c.Contents)
}

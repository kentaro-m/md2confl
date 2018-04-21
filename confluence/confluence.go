package confluence

import (
	bfconfluence "github.com/kentaro-m/blackfriday-confluence"
)

// Confluence is a type that implements the interface for confluence wiki output.
type Confluence struct {
	contents string
}

// Convert renders confluence wiki text from markdown text.
func (c *Confluence) Convert(input []byte) {
	output := bfconfluence.Run(input)
	c.contents = string(output)
}

// Contents returns confluence wiki output.
func (c *Confluence) Contents() string {
	return c.contents
}

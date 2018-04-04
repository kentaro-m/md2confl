package confluence

import (
	"fmt"

	bfconfluence "github.com/kentaro-m/blackfriday-confluence"
)

// Confluence is a type that implements the interface for confluence wiki output.
type Confluence struct {
	Contents string
}

// Convert renders confluence wiki text from markdown text.
func (c *Confluence) Convert(input []byte) {
	output := bfconfluence.Run(input)
	c.Contents = string(output)
}

// Preview displays confluence wiki document.
func (c *Confluence) Preview() {
	fmt.Println(c.Contents)
}

package bitbar

import (
	"bytes"
	"fmt"

	"github.com/schigh/servicemonger/ui"
)

type Builder interface {
	// EOL will end the current line.
	EOL()
	// Strf writes a formatted string with optional
	// arguments to the current line
	Strf(format string, args ...interface{}) Builder
	// Pipe adds a separator marking the boundary between
	// a line item's title and its parameters
	Pipe() Builder
	// Nest will add a level of nesting to the current line,
	// defined by the depth.  This function should be called
	// at the beginning of the target line.
	Nest(depth int) Builder
	// Command will render a bash command
	Command(c *Command) Builder
	// FOnt will render a font
	Font(font ui.Font) Builder
	// Typeface will render a typeface name within a font
	Typeface(name string) Builder
	// Size will render the size of the current line
	Size(size int) Builder
	// Color will render the color of the current line
	Color(color ui.Color) Builder
	// Href will render the link for the current line
	Href(href string) Builder
	// Image will render an image on the current line
	Image(img string) Builder
	// TemplateImage will render the main status bar image
	TemplateImage(img string) Builder
	// Bytes will return the buffered bytes to be written out
	Bytes() []byte
}

type builder struct {
	bytes.Buffer
}

func NewBuilder() *builder {
	return &builder{}
}

func (b *builder) EOL() {
	_ = b.WriteByte('\n')
}

func (b *builder) Strf(format string, args ...interface{}) Builder {
	_, _ = b.Write([]byte(fmt.Sprintf(format, args...)))
	return b
}

func (b *builder) Pipe() Builder {
	_, _ = b.WriteString(" |")
	return b
}

func (b *builder) Nest(depth int) Builder {
	_, _ = b.Write(bytes.Repeat([]byte{'-', '-'}, depth))
	return b
}

func (b *builder) Command(c *Command) Builder {
	if c != nil {
		_, _ = b.Write([]byte(fmt.Sprintf(` bash="%s"`, c.Exec)))
		for i, p := range c.Params {
			_, _ = b.Write([]byte(fmt.Sprintf(` param%d="%s"`, i+1, p)))
		}
	}

	return b
}

func (b *builder) Font(font ui.Font) Builder {
	return b.Typeface(font.Name).Size(font.Size).Color(ui.ColorWithID(font.ColorID))
}

func (b *builder) Typeface(name string) Builder {
	if name != "" {
		_, _ = b.Write([]byte(fmt.Sprintf(" font=%s", name)))
	}
	return b
}

func (b *builder) Size(size int) Builder {
	if size > 0 {
		_, _ = b.Write([]byte(fmt.Sprintf(" size=%d", size)))
	}
	return b
}

func (b *builder) Color(color ui.Color) Builder {
	if color != ui.DefaultColor {
		_, _ = b.Write([]byte(fmt.Sprintf(" color=%s", color)))
	}
	return b
}

func (b *builder) Href(href string) Builder {
	if href != "" {
		_, _ = b.Write([]byte(fmt.Sprintf(" href=%s", href)))
	}
	return b
}

func (b *builder) Image(image string) Builder {
	if image != "" {
		_, _ = b.Write([]byte(fmt.Sprintf(" image=%s", image)))
	}
	return b
}

func (b *builder) TemplateImage(image string) Builder {
	_, _ = b.Write([]byte(fmt.Sprintf(" templateImage=%s", image)))
	return b
}

func (b *builder) Bytes() []byte {
	return b.Buffer.Bytes()
}

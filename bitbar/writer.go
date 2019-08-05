package bitbar

import "fmt"

type Writer interface {
	Write(c *Config, b Builder) error
}

type writer struct{}

func NewWriter() *writer {
	return &writer{}
}

func (w *writer) Write(c *Config, b Builder) error {
	switch {
	case c.StatusBarImage != "":
	//TODO output image
	case c.StatusBarEmoji != "":
		b.Strf(":%s:", c.StatusBarEmoji).EOL()
	default:
		b.Strf(":japanese_ogre:").EOL()
	}

	b.Strf("---").EOL()
	for _, p := range c.Services {
		b.Strf(p.Title).Pipe().Font(c.Fonts["header"]).EOL()
		for _, l := range p.Links {
			b.Nest(1)
			switch {
			case l.Image != "":
			//TODO output image
			case l.Emoji != "":
				b.Strf(":%s: ", l.Emoji)
			}

			// ensure the default font
			if l.FontID == "" {
				l.FontID = "default"
			}
			b.Strf(l.Title).Pipe().Href(l.Href).Font(c.Fonts[l.FontID]).EOL()
		}
	}

	bb := b.Bytes()
	fmt.Printf("%s", bb)
	return nil
}

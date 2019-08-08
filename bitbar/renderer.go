package bitbar

import (
	"fmt"

	"github.com/schigh/servicemonger/ui"
)

type Renderer interface {
	ServiceRenderer
	LinkRenderer
	// RenderBitBar is the top-level rendering entrypoint
	RenderBitBar(*Config, Builder) error
}

type renderer struct{}

// DefaultRenderer returns the default rendering engine for servicemonger
func DefaultRenderer() *renderer {
	return &renderer{}
}

func (r *renderer) RenderBitBar(c *Config, b Builder) error {
	switch {
	case c.StatusBarImage != "":
		b.Pipe().TemplateImage(c.StatusBarImage).EOL()
	case c.StatusBarEmoji != "":
		b.Strf(":%s:", c.StatusBarEmoji).EOL()
	default:
		b.Pipe().TemplateImage(ui.DefaultStatusBarImage()).EOL()
	}

	b.Strf("---").EOL()
	opts := ServiceRenderOptions{
		Builder: b,
		Config:  c,
		Depth:   0,
	}
	for _, s := range c.Services {
		opts.Service = s
		if err := r.RenderService(opts); err != nil {
			return err
		}
	}

	bb := b.Bytes()
	fmt.Printf("%s", bb)
	return nil
}

func (r *renderer) RenderService(opts ServiceRenderOptions) error {
	b := opts.Builder
	s := opts.Service
	cfg := opts.Config
	if opts.Depth > 0 {
		b.Nest(opts.Depth)
	}
	b.Strf(s.Title).Pipe().Font(cfg.Fonts["header"]).EOL()
	lOpts := LinkRenderOptions{
		Builder: b,
		Config:  cfg,
		Depth:   opts.Depth + 1,
	}
	for _, l := range s.Links {
		lOpts.Link = l
		if err := r.RenderLink(lOpts); err != nil {
			return err
		}
	}
	return nil
}

func (r *renderer) RenderLink(opts LinkRenderOptions) error {
	b := opts.Builder
	l := opts.Link
	cfg := opts.Config
	if opts.Depth > 0 {
		b.Nest(opts.Depth)
	}

	if l.Emoji != "" {
		b.Strf(":%s: ", l.Emoji)
	}

	if l.Font == "" {
		l.Font = "default"
	}

	b.Strf(l.Title).
		Pipe().
		Href(l.Href).
		Font(cfg.Fonts[l.Font]).
		Image(l.Image).
		Command(l.Command).
		EOL()
	return nil
}

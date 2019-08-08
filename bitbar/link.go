package bitbar

type LinkRenderOptions struct {
	Link    Link
	Builder Builder
	Config  *Config
	Depth   int
}

type LinkRenderer interface {
	RenderLink(LinkRenderOptions) error
}

type Link struct {
	Title   string   `yaml:"title"`
	Href    string   `yaml:"href"`
	Command *Command `yaml:"command,omitempty"`
	Emoji   string   `yaml:"emoji"`
	Image   string   `yaml:"image"`
	Font    string   `yaml:"font"`
}

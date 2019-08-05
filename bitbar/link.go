package bitbar

type Link struct {
	Title   string   `yaml:"title"`
	Href    string   `yaml:"href"`
	Command *Command `yaml:"command,omitempty"`
	Emoji   string   `yaml:"emoji"`
	Image   string   `yaml:"image"`
	FontID  string   `yaml:"font"`
}

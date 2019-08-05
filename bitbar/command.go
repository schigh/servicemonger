package bitbar

type Command struct {
	Title  string   `yaml:"title"`
	Exec   string   `yaml:"exec"`
	Params []string `yaml:"params"`
}

package bitbar

type Command struct {
	Exec   string   `yaml:"exec"`
	Params []string `yaml:"params"`
}

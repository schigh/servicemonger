package bitbar

type Service struct {
	Title string `yaml:"title"`
	Links []Link `yaml:"links"`
	FontID string `yaml:"font"`
}

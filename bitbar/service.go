package bitbar

// ServiceRenderer knows how to render Service structs
type ServiceRenderOptions struct {
	Service Service
	Builder Builder
	Config  *Config
	Depth   int
}
type ServiceRenderer interface {
	RenderService(ServiceRenderOptions) error
}

// Service defines what is loosely described as a "service" that has submenus and such
type Service struct {
	Title    string    `yaml:"title"`
	Links    []Link    `yaml:"links"`
	FontID   string    `yaml:"font"`
	Services []Service `yaml:"services"`
}

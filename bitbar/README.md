# bitbar
--
    import "."


## Usage

#### func  DefaultRenderer

```go
func DefaultRenderer() *renderer
```
DefaultRenderer returns the default rendering engine for servicemonger

#### func  NewBuilder

```go
func NewBuilder() *builder
```

#### type Builder

```go
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
```


#### type Command

```go
type Command struct {
	Exec   string   `yaml:"exec"`
	Params []string `yaml:"params"`
}
```


#### type Config

```go
type Config struct {
	StatusBarImage string                     `yaml:"status_bar_image"`
	StatusBarEmoji string                     `yaml:"status_bar_emoji"`
	Fonts          map[string]ui.Font         `yaml:"fonts"`
	Colors         map[string]ui.Color        `yaml:"colors"`
	Credentials    map[string]auth.Credential `yaml:"credentials"`
	Commands       map[string]Command         `yaml:"commands"`
	Images         map[string]string          `yaml:"images"`
	Services       []Service                  `yaml:"services"`
}
```


#### func  LoadConfig

```go
func LoadConfig(cfgPath string) (*Config, error)
```

#### type Link

```go
type Link struct {
	Title   string   `yaml:"title"`
	Href    string   `yaml:"href"`
	Command *Command `yaml:"command,omitempty"`
	Emoji   string   `yaml:"emoji"`
	Image   string   `yaml:"image"`
	Font    string   `yaml:"font"`
}
```


#### type LinkRenderOptions

```go
type LinkRenderOptions struct {
	Link    Link
	Builder Builder
	Config  *Config
	Depth   int
}
```


#### type LinkRenderer

```go
type LinkRenderer interface {
	RenderLink(LinkRenderOptions) error
}
```


#### type Renderer

```go
type Renderer interface {
	ServiceRenderer
	LinkRenderer
	// RenderBitBar is the top-level rendering entrypoint
	RenderBitBar(*Config, Builder) error
}
```


#### type Service

```go
type Service struct {
	Title    string    `yaml:"title"`
	Links    []Link    `yaml:"links"`
	FontID   string    `yaml:"font"`
	Services []Service `yaml:"services"`
}
```

Service defines what is loosely described as a "service" that has submenus and
such

#### type ServiceRenderOptions

```go
type ServiceRenderOptions struct {
	Service Service
	Builder Builder
	Config  *Config
	Depth   int
}
```

ServiceRenderer knows how to render Service structs

#### type ServiceRenderer

```go
type ServiceRenderer interface {
	RenderService(ServiceRenderOptions) error
}
```

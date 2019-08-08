# ui
--
    import "."


## Usage

```go
const (
	DefaultColor  = Color("") // this is a no-op ... app will use the system color
	DarkColor     = Color("#313131")
	LightColor    = Color("#EAEAEA")
	InactiveColor = Color("#909090")
	WarnColor     = Color("#F7BB08")
	AlertColor    = Color("#F70808")
	OKColor       = Color("#42E662")
)
```

#### func  DefaultStatusBarImage

```go
func DefaultStatusBarImage() string
```

#### type Color

```go
type Color string
```


#### func  ColorWithID

```go
func ColorWithID(id string) Color
```

#### type Font

```go
type Font struct {
	Name    string `yaml:"name"`
	Size    int    `yaml:"size"`
	ColorID string `yaml:"color"`
}
```

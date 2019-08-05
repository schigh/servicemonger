package ui

type Color string

var colorCache map[string]Color

const (
	DefaultColor  = Color("") // this is a no-op ... app will use the system color
	DarkColor     = Color("#313131")
	LightColor    = Color("#EAEAEA")
	InactiveColor = Color("#909090")
	WarnColor     = Color("#F7BB08")
	AlertColor    = Color("#F70808")
	OKColor       = Color("#42E662")
)

func ColorWithID(id string) Color {
	if color, ok := colorCache[id]; ok {
		return color
	}

	switch id {
	case "dark":
		return DarkColor
	case "light":
		return LightColor
	case "inactive":
		return InactiveColor
	case "warn", "warning":
		return WarnColor
	case "alert", "error", "failure":
		return AlertColor
	case "ok", "success":
		return OKColor
	}

	return DefaultColor
}

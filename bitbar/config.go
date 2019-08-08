package bitbar

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	"github.com/schigh/servicemonger/auth"
	"github.com/schigh/servicemonger/ui"
)

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

func LoadConfig(cfgPath string) (*Config, error) {
	if cfgPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		cfgPath = path.Join(home, ".servicemonger.yml")
	}

	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var s Config
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

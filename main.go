package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/schigh/servicemonger/bitbar"
)

var (
	cfgPath string
)

func main() {
	flag.StringVar(&cfgPath, "settings", "", "absolute path to settings")
	flag.Parse()

	settings, err := bitbar.LoadConfig(cfgPath)
	if err != nil {
		fmt.Printf("unable to load settings: %v\n", err)
		os.Exit(1)
	}

	b := bitbar.NewBuilder()
	w := bitbar.NewWriter()
	_ = w.Write(settings, b)
}

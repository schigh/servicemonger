package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/schigh/servicemonger/b64"
	"github.com/schigh/servicemonger/bitbar"
)

var (
	cfgPath      string
	filePath     string
	shouldBase64 bool
)

func main() {
	flag.StringVar(&cfgPath, "settings", "", "absolute path to settings")
	flag.BoolVar(&shouldBase64, "b64", false, "base64-encode the provided image path")
	flag.StringVar(&filePath, "file", "", "path to file in context")
	flag.Parse()

	// bypass render
	if shouldBase64 {
		imgStr, err := b64.RenderBase64ForImage(filePath)
		if err != nil {
			fmt.Printf("unable to base64-encode image: %v\n", err)
			os.Exit(1)
		}
		println(imgStr)
		os.Exit(0)
	}

	settings, err := bitbar.LoadConfig(cfgPath)
	if err != nil {
		fmt.Printf("unable to load settings: %v\n", err)
		os.Exit(1)
	}

	b := bitbar.NewBuilder()
	r := bitbar.DefaultRenderer()
	_ = r.RenderBitBar(settings, b)
}

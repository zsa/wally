package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/zsa/wally/state"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	_state := &state.State{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Wally",
		Width:            630,
		Height:           520,
		MinWidth:           630,
        	MinHeight:          520,
        	MaxWidth:           630,
        	MaxHeight:          520,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        _state.Init,
		DisableResize:    true,
		Bind: []interface{}{
			_state,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	defer _state.Teardown()
}

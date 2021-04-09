package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
	"github.com/zsa/wally/wally"
)

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:     630,
		Height:    520,
		Resizable: false,
		Title:     "Wally",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})
	state := wally.NewState(wally.Probing, "")
	app.Bind(state)
	app.Run()
}

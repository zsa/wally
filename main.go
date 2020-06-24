package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     630,
		Height:    520,
		Resizable: false,
		Title:     "Wally",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})
	state := NewState(0, "")
	app.Bind(state)
	app.Run()
}

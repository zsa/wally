package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"github.com/zsa/wally/wally"
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
	state := wally.NewState(wally.Probing, "")
	app.Bind(state)
	app.Run()
}

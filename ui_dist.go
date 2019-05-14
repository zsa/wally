// +build dist

package main

import (
	"github.com/fdidron/webview"
	"wally/wally"
)

// Init returns a configured and ready to use webview.
// Used with the 'dist' build tag. All ui assets are embedded into the binary using go-bindata
func Init() (wv webview.WebView) {
	State := wally.State{Step: 0}

	html := `
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Wally</title>
</head>
<body>
</html>
	`

	js := string(MustAsset("index.dist.js"))

	w := webview.New(webview.Settings{
		Debug:     false,
		Width:     630,
		Height:    440,
		Title:     "Wally",
		Resizable: false,
		URL:       `data:text/html,` + html,
		ExternalInvokeCallback: func(w webview.WebView, command string) {
			handleRPC(w, command, &State)
		},
	})

	w.Dispatch(func() {
		w.Bind("state", &State)
		w.Eval(js)
	})

	State.Log("info", "Application UI / State initialized")

	return w

}

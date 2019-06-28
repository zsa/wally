// +build dev

package main

import (
	"github.com/fdidron/webview"
	"github.com/zsa/wally/wally"
)

// Init returns a configured and ready to use webview.
// Used with the 'dev' build tag. All ui assets are expected to be served locally on port 8080. A local dev server can be run from the ui folder running yarn dev.
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

	js := `
(function(){
	var n=document.createElement('script');
	n.setAttribute('type', 'text/javascript');
	n.setAttribute('src', 'http://localhost:8080/index.dist.js');
	document.body.appendChild(n);
})()
`

	w := webview.New(webview.Settings{
		Debug:     true,
		Width:     630,
		Height:    500,
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

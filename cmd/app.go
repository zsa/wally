package cmd

import (
	_ "embed"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails"
	"github.com/zsa/wally/wally"
)

var stylesheet = path.Join("frontend/build/static/css/main.css")
var javascript = path.Join("frontend/build/static/js/main.js")

func runAppCmd(cmd *cobra.Command, args []string) {

	javascript = viper.GetString("javascript")
	stylesheet = viper.GetString("stylesheet")

	js, err := os.ReadFile(javascript)
	if err != nil {
		log.Fatal(err)
	}

	css, err := os.ReadFile(stylesheet)
	if err != nil {
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     630,
		Height:    520,
		Resizable: false,
		Title:     "Wally",
		Colour:    "#131313",
		JS:        string(js),
		CSS:       string(css),
	})
	state, err := wally.NewState(wally.Probing, "")
	if err != nil {
		panic(err)
	}
	app.Bind(state)
	app.Run()
	return
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Launch Wally as an application",
	Run:   runAppCmd,
}

func init() {
	rootCmd.AddCommand(appCmd)

	appCmd.Flags().StringVarP(&javascript, "javascript", "j", javascript, "javascript file")
	appCmd.Flags().StringVarP(&stylesheet, "stylesheet", "s", stylesheet, "css file")

	viper.SetDefault("javascript", javascript)
	viper.SetDefault("stylesheet", stylesheet)

}

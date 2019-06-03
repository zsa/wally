package main

import (
	"fmt"
	"github.com/caarlos0/spin"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"path/filepath"
	"time"
	"github.com/zsa/wally/wally"
)

func main() {
	var args = os.Args[1:]
	if len(args) != 1 {
		fmt.Println(Blue("Usage: wally-cli <firmware file>"))
		os.Exit(1)
	}

	path := args[0]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(Red("The file path you specified does not exist"))
		os.Exit(1)
	}

	extension := filepath.Ext(path)
	if extension != ".bin" && extension != ".hex" {
		fmt.Println(Red("The file you specified should be a"), Red(Underline(".hex")), Red("file (ErgoDox EZ) or a"), Red(Underline(".bin")), Red("file (Planck EZ)"))
		os.Exit(1)

	}

	s := wally.State{Step: 3}
	if extension == ".bin" {
		go wally.DFUFlash(path, &s)
	}
	if extension == ".hex" {
		go wally.TeensyFlash(path, &s)
	}

	spinner := spin.New("%s Press the reset button of your keyboard.")
	spinner.Start()
	spinnerStopped := false

	var progress *pb.ProgressBar
	progressStarted := false

	for s.Step != 5 {
		time.Sleep(500 * time.Millisecond)
		if s.FlashProgress.Step > 0 {
			if spinnerStopped == false {
				spinner.Stop()
				spinnerStopped = true
			}
			if progressStarted == false {
				progressStarted = true
				progress = pb.StartNew(s.FlashProgress.Total)
			}
			progress.Set(s.FlashProgress.Sent)
		}
	}
	progress.Finish()
	fmt.Println(Green("Your keyboard was successfully flashed and rebooted. Enjoy the new firmware!"))
}

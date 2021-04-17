package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/caarlos0/spin"
	"github.com/spf13/cobra"
	"github.com/zsa/wally/wally"
	"gopkg.in/cheggaaa/pb.v1"
)

func runFlashCmd(cmd *cobra.Command, args []string) {
	firmwarePath := args[0]

	// s := wally.State{Step: 0, Total: 0, Sent: 0}
	state, err := wally.NewState(wally.Probing, firmwarePath)
	if err != nil {
		log.Fatalln(err)
	}

	ext := filepath.Ext(firmwarePath)
	if ext == ".bin" {
		go wally.DFUFlash(state)
	} else if ext == ".hex" {
		go wally.TeensyFlash(state)
	} else {
		log.Fatalf("unsupported file extension: %v", ext)
	}

	var progress *pb.ProgressBar
	progressStarted := false

	spinner := spin.New("%s Press the reset button of your keyboard.")
	spinner.Start()
	spinnerStopped := false

	state.Step = wally.Waiting

	// Loop until we have fully flashed the firmware file
	for state.Step != wally.Complete {
		// Pause before probing once more
		if state.Step == wally.Waiting {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		// If we are no longer probing, and we have not begun tracking
		// progress, create a new progress bar to print to the console
		if progressStarted == false {
			progress = pb.StartNew(state.FlashProgress.Total)
			progressStarted = true
		}
		// If we are no longer probing, and we have not stopped the "pinwheel",
		// spinning as we idle, then now is the time to do so
		if spinnerStopped == false {
			spinner.Stop()
			spinnerStopped = true
		}
		// Update the progress bar with our current status
		progress.Set(state.FlashProgress.Sent)
	}
	progress.Finish()
	fmt.Println("Your keyboard was successfully flashed and rebooted. Enjoy the new firmware!")
}

var flashCmd = &cobra.Command{
	Use:     "flash",
	Short:   "Flash a firmware file directly to the keyboard",
	Example: "flash {firmware.hex | firmware.bin}",
	Args:    cobra.ExactArgs(1),
	Run:     runFlashCmd,
}

var configFile string

// The ConsoleFlash function

func init() {
	rootCmd.AddCommand(flashCmd)
}

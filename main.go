package main

import "os"

func main() {
	var filePath string

	args := os.Args[1:]

	if len(args) > 0 {
		filePath = args[0]
	}

	w := Init(filePath)
	defer w.Exit()

	w.Run()
}

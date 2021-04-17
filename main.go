package main

import (
	"log"

	"github.com/zsa/wally/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

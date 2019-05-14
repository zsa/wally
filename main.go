package main

func main() {

	w := Init()
	defer w.Exit()

	w.Run()
}

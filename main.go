package main

import (
	"os"
)

func main() {
	if len(os.Args) == 1 {
		printUsage()
	}

	fileName := os.Args[1]
	In := readFile(fileName)

	os.WriteFile(In.Name+".html", *In.Data, 0644)
}

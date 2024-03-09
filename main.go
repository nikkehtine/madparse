package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		printUsage()
	}

	fileName := os.Args[1]
	In := readFile(fileName)
	Out := &File{
		Name:      In.Name,
		Extension: "html",
		Data:      In.Data,
	}

	fmt.Print(string(In.Data))
	writeFile(Out)
}

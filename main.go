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
		name:      In.name,
		extension: "html",
		data:      In.data,
	}

	fmt.Print(string(In.data))
	writeFile(Out)
}

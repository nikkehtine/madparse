package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("madparse filename --flags")
	os.Exit(1)
}

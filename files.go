package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	name      string
	extension string
	data      []byte
}

func readFile(fileName string) *File {
	file := new(File)
	body, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		log.Fatalln("File does not exist")
	} else if err != nil {
		log.Fatalln(err)
	} else {
		file.extension = filepath.Ext(fileName)
		file.name = strings.TrimSuffix(filepath.Base(fileName), file.extension)
		file.data = body
	}

	return file
}

func writeFile(Out *File) error {
	fileName := Out.name + "." + Out.extension
	err := os.WriteFile(fileName, Out.data, 0644)
	return err
}

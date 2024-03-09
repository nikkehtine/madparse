package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name      string
	Extension string
	Data      []byte
}

func readFile(fileName string) *File {
	file := new(File)
	body, err := os.ReadFile(fileName)
	if os.IsNotExist(err) {
		log.Fatalln("File does not exist")
	} else if err != nil {
		log.Fatalln(err)
	} else {
		file.Extension = filepath.Ext(fileName)
		file.Name = strings.TrimSuffix(filepath.Base(fileName), file.Extension)
		file.Data = body
	}

	return file
}

func writeFile(Out *File) error {
	fileName := Out.Name + "." + Out.Extension
	err := os.WriteFile(fileName, Out.Data, 0644)
	return err
}

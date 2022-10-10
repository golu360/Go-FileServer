package utils

import (
	"fmt"
	"log"
	"os"
)

type FileUtils struct {
	DirName string
}

func (f *FileUtils) GetKeys() []string {
	var directories []string
	fs, err := os.ReadDir("data")
	if err != nil {
		fmt.Println(err)
		return directories
	}
	for _, f := range fs {
		if f.Type().IsDir() {
			directories = append(directories, f.Name())
		}
	}
	return directories
}

func (f *FileUtils) CreateKey(KeyName string) bool {
	if _, err := os.Stat("data/" + KeyName); !os.IsNotExist(err) {
		log.Println(err)
		return false
	}
	err := os.Mkdir(f.DirName+"/"+KeyName, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

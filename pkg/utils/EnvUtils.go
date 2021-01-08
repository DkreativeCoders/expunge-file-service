package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetPathToFileCleanerJson() string{
	workingDirectorypath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("working directory", workingDirectorypath)
	configPath := os.Getenv("CONFIG.PATH")
	fullPathToFileCleanerJson := filepath.Join(workingDirectorypath, configPath)
	fmt.Println("Full Path directory", fullPathToFileCleanerJson)
	return fullPathToFileCleanerJson
}

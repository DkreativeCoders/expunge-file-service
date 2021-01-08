package utils

import (
	"fmt"
	"log"
	"os"
)

func GetPathToFileCleanerJson() string{
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("working directory", path)

	configPath := os.Getenv("CONFIG.PATH")
	fullPathToFileCleanerJson := path + configPath
	fmt.Println("Full Path directory", fullPathToFileCleanerJson)
	return fullPathToFileCleanerJson
}

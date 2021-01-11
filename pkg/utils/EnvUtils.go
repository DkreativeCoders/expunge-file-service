package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPathToFileCleanerJson() string{
	workingDirectoryPath, err := os.Getwd()
	if err != nil {
	}
	fmt.Println("working directory", workingDirectoryPath)
	configPath := os.Getenv("CONFIG.PATH")
	fullPathToFileCleanerJson := filepath.Join(workingDirectoryPath, configPath)
	fmt.Println("Full Path directory", fullPathToFileCleanerJson)
	return fullPathToFileCleanerJson
}
